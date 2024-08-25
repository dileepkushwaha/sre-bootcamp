# Stage 1: Build the application
FROM golang:1.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /sre-bootcamp ./cmd

# Stage 2: Create a small image to run the application
FROM alpine:latest

# Install necessary packages
RUN apk --no-cache add ca-certificates postgresql-client

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /sre-bootcamp .

# Copy the .env file into the container
COPY .env /root/.env

# Copy the startup script
COPY start.sh /root/

# Give execute permission to the startup script
RUN chmod +x /root/start.sh

# Copy migration files and run-migrations.sh script
COPY migrations/ /migrations/
COPY run-migrations.sh /root/run-migrations.sh

# Give execute permission to the migration script
RUN chmod +x /root/run-migrations.sh

# Set an environment variable for the port
ENV APP_PORT=8080

# Copy the wrapper script
COPY start-wrapper.sh /root/start-wrapper.sh

# Make the wrapper script executable
RUN chmod +x /root/start-wrapper.sh

# Expose ports for API
EXPOSE 8080
EXPOSE 8081

# Entry point to run migrations and then start the application
CMD ["/bin/sh", "-c", "/root/run-migrations.sh && ./sre-bootcamp --port $APP_PORT && /root/start.sh && /root/start-wrapper.sh"]
