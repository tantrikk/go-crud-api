.PHONY: help build run stop clean logs restart status test

# Default target
help:
	@echo "Available commands:"
	@echo "  make build    - Build the Docker image"
	@echo "  make run      - Build and run the container"
	@echo "  make stop     - Stop the running container"
	@echo "  make clean    - Stop and remove container and image"
	@echo "  make logs     - Show container logs"
	@echo "  make restart  - Restart the container"
	@echo "  make status   - Show container status"
	@echo "  make test     - Test the API endpoints"

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
test:
	@echo "Testing API endpoints..."
	@echo "\n1. Creating user..."
	@curl -X POST http://localhost:8080/users \
		-H "Content-Type: application/json" \
		-d '{"name":"Test User","email":"test@example.com","password":"test123"}' \
		-s | jq . || echo "Failed to create user"
	@echo "\n2. API is working!"