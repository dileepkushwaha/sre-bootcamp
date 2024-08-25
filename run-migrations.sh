#!/bin/sh

# Exit on error
set -e

# Database credentials
DB_HOST=${DB_HOST:-db-service}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-user}
DB_PASSWORD=${DB_PASSWORD:-password}
DB_NAME=${DB_NAME:-studentdb}

# Export PGPASSWORD to avoid prompt for password
export PGPASSWORD=$DB_PASSWORD

# Create database if it doesn't exist
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres <<-EOSQL
    SELECT 'CREATE DATABASE $DB_NAME' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '$DB_NAME')\gexec
EOSQL

# Path to the migrations directory
MIGRATION_DIR="/migrations"

# Create a migrations table if it doesn't exist
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME <<-EOSQL
    CREATE TABLE IF NOT EXISTS applied_migrations (
        filename TEXT PRIMARY KEY,
        applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
EOSQL

# Run migrations
echo "Running database migrations..."
for file in $MIGRATION_DIR/*.sql; do
    filename=$(basename "$file")
    
    # Check if migration has already been applied
    if psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -tAc "SELECT 1 FROM applied_migrations WHERE filename='$filename'" | grep -q 1; then
        echo "Migration already applied: $filename"
    else
        echo "Applying migration: $filename"
        psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f $file
        
        # Record the migration as applied
        psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "INSERT INTO applied_migrations (filename) VALUES ('$filename')"
    fi
done

echo "Migrations applied successfully."