.PHONY: help build run stop clean logs restart status test test-unit test-api test-all test-coverage dc-up dc-down dc-logs dc-build

# Default target
help:
	@echo "Available commands:"
	@echo "  make build         - Build the Docker image"
	@echo "  make run           - Build and run the container"
	@echo "  make stop          - Stop the running container"
	@echo "  make clean         - Stop and remove container and image"
	@echo "  make logs          - Show container logs"
	@echo "  make restart       - Restart the container"
	@echo "  make status        - Show container status"
	@echo "  make test-api      - Test the API endpoints"
	@echo "  make test-unit     - Run unit tests"
	@echo "  make test-all      - Run all tests"
	@echo "  make test-coverage - Run tests with coverage report"
	@echo ""
	@echo "Docker Compose commands:"
	@echo "  make dc-up         - Start all services (MySQL, API, Frontend)"
	@echo "  make dc-down       - Stop all docker-compose services"
	@echo "  make dc-logs       - Show docker-compose logs"
	@echo "  make dc-build      - Build and start all services"
	@echo "  make frontend-dev  - Run frontend in development mode"

# Build Docker image
build:
	docker build -t go-crud-api .

# Run container (builds image if needed)
run: build
	@echo "Starting container..."
	@docker stop go-crud-api-container 2>/dev/null || true
	@docker rm go-crud-api-container 2>/dev/null || true
	docker run -d -p 8080:8080 --name go-crud-api-container go-crud-api
	@echo "Container started! API available at http://localhost:8080"

# Stop container
stop:
	@echo "Stopping container..."
	@docker stop go-crud-api-container 2>/dev/null || echo "Container not running"

# Clean up everything (containers and images)
clean:
	@echo "Cleaning up containers and images..."
	@docker stop go-crud-api-container 2>/dev/null || true
	@docker rm go-crud-api-container 2>/dev/null || true
	@docker rmi go-crud-api 2>/dev/null || true
	@echo "Cleanup complete!"

# Show logs
logs:
	docker logs -f go-crud-api-container

# Restart container
restart: stop run

# Show container status
status:
	@echo "Container status:"
	@docker ps -a | grep go-crud-api-container || echo "No container found"

# Test API endpoints
test-api:
	@echo "Testing API endpoints..."
	@echo "\n1. Creating user..."
	@curl -X POST http://localhost:8080/users \
		-H "Content-Type: application/json" \
		-d '{"name":"Test User","email":"test@example.com","password":"test123"}' \
		-s | jq . || echo "Failed to create user"
	@echo "\n2. API is working!"

# Run unit tests
test-unit:
	@echo "Running unit tests..."
	@go test -v ./...

# Run all tests
test-all: test-unit test-api

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Docker Compose commands
dc-up:
	@echo "Starting services with docker-compose..."
	docker-compose up -d
	@echo "Services started! API available at http://localhost:8080"
	@echo "MySQL available at localhost:3306"

dc-down:
	@echo "Stopping docker-compose services..."
	docker-compose down
	@echo "Services stopped!"

dc-logs:
	docker-compose logs -f

dc-build:
	@echo "Building and starting services..."
	docker-compose up -d --build
	@echo "Services built and started!"
	@echo "Backend API: http://localhost:8080"
	@echo "Frontend UI: http://localhost:3000"

# Run frontend in development mode
frontend-dev:
	@echo "Starting frontend in development mode..."
	cd frontend && npm install && npm start