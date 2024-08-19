# Define variables
APP_NAME = sre-bootcamp
GO_FILES = $(shell find . -name '*.go')
BUILD_DIR = bin
BUILD_TARGET = $(BUILD_DIR)/$(APP_NAME)

# Default target
all: build

# Build the application
build: $(GO_FILES)
	@echo "Building the application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_TARGET) ./cmd

# Run the application
run: build
	@echo "Running the application..."
	@$(BUILD_TARGET)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Apply database migrations
migrate:
	@echo "Applying database migrations..."
	@go run ./migrations

# Check the health of the application
healthcheck:
	@echo "Checking health..."
	@curl -s http://localhost:8080/healthcheck

# Docker variables
IMAGE_NAME = sre-bootcamp
IMAGE_TAG = 1.0.0
CONTAINER_NAME = sre-bootcamp
DB_CONTAINER_NAME = my_database

# Start the DB container
db-start:
	@echo "Starting the DB container..."
	docker-compose up -d db

# Run DB migrations
db-migrate:
	@echo "Running DB migrations..."
	# Add your migration command here
	# Example: docker run --rm -it my-migration-tool

# Build the Docker image for the REST API
build-api:
	@echo "Building Docker image for the API..."
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

# Run the REST API container
run-api: db-start db-migrate build-api
	@echo "Running the API container..."
	docker-compose up -d api

# Stop and remove all containers
stop-all:
	@echo "Stopping all containers..."
	docker-compose down

# Docker variables
IMAGE_NAME = sre-bootcamp
IMAGE_TAG = 1.0.0
CONTAINER_NAME = sre-bootcamp
DB_CONTAINER_NAME = my_database

# Linting
lint:
	@echo "Running code linting..."
	golangci-lint run

# Build the Docker image for the REST API
build-api:
	@echo "Building Docker image for the API..."
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

# Push the Docker image to the registry
push-image:
	@echo "Pushing Docker image to the registry..."
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) ghcr.io/$(shell git config user.name)/$(IMAGE_NAME):$(IMAGE_TAG)
	docker push ghcr.io/$(shell git config user.name)/$(IMAGE_NAME):$(IMAGE_TAG)

# Build and push the Docker image
build-and-push: build-api push-image

.PHONY: lint build-api push-image build-and-push
.PHONY: db-start db-migrate build-api run-api stop-all

.PHONY: all build run clean test migrate healthcheck
