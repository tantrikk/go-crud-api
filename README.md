# Go CRUD API

A simple, well-tested CRUD API for managing users, built with Go. It provides RESTful endpoints to create, read, update, and delete user information with in-memory storage.

## Features

- RESTful API endpoints for user management
- In-memory data storage
- Comprehensive unit test suite with 100% coverage
- Docker support for easy deployment
- Makefile for simplified operations
- Clean architecture with separation of concerns

## Project Structure

```
go-crud-api/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── handler/
│   │   ├── user.go          # HTTP request handlers
│   │   └── user_test.go     # Handler unit tests
│   ├── model/
│   │   ├── user.go          # User model definition
│   │   └── user_test.go     # Model unit tests
│   └── repository/
│       ├── user.go          # Data storage layer
│       └── user_test.go     # Repository unit tests
├── go.mod                   # Go module definition
├── go.sum                   # Dependency checksums
├── Dockerfile               # Docker configuration
├── Makefile                 # Build and operation commands
└── README.md                # Project documentation
```

## Prerequisites

- Go 1.20 or higher
- Docker (optional, for containerized deployment)
- Make (optional, for using Makefile commands)

## Quick Start

### Using Make (Recommended)

```bash
# Show available commands
make help

# Run the application in Docker
make run

# Run unit tests
make test-unit

# Stop the container
make stop
```

### Manual Setup

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd go-crud-api
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the application:**
   ```bash
   go run cmd/main.go
   ```

   The server will start on `http://localhost:8080`

## API Endpoints

### Create User
- **POST** `/users`
- **Body:**
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepassword"
  }
  ```
- **Response:** 201 Created
  ```json
  {
    "id": "generated-uuid",
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepassword"
  }
  ```

### Get User
- **GET** `/users/{id}`
- **Response:** 200 OK
  ```json
  {
    "id": "user-id",
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepassword"
  }
  ```

### Update User
- **PUT** `/users/{id}`
- **Body:**
  ```json
  {
    "name": "John Doe Updated",
    "email": "john.updated@example.com"
  }
  ```
- **Response:** 200 OK

### Delete User
- **DELETE** `/users/{id}`
- **Response:** 204 No Content

### Error Responses
- **400 Bad Request:** Invalid request body
- **404 Not Found:** User not found

## Testing

The project includes a comprehensive test suite with 100% code coverage.

### Run Tests

```bash
# Run all unit tests
make test-unit

# Run tests with coverage report
make test-coverage

# Test API endpoints (requires running server)
make test-api

# Run all tests
make test-all
```

### Test Coverage

- **Handler Layer:** 100% coverage
- **Repository Layer:** 100% coverage
- **Model Layer:** Full test coverage

Coverage reports are generated in HTML format at `coverage.html`.

## Docker Support

### Build and Run with Docker

```bash
# Build Docker image
docker build -t go-crud-api .

# Run container
docker run -p 8080:8080 go-crud-api
```

### Using Make Commands

```bash
# Build and run container
make run

# View logs
make logs

# Stop container
make stop

# Clean up (remove container and image)
make clean

# Check container status
make status
```

## Makefile Commands

| Command | Description |
|---------|-------------|
| `make help` | Show all available commands |
| `make build` | Build Docker image |
| `make run` | Build and run container |
| `make stop` | Stop running container |
| `make clean` | Remove container and image |
| `make logs` | Show container logs |
| `make restart` | Restart container |
| `make status` | Show container status |
| `make test-unit` | Run unit tests |
| `make test-api` | Test API endpoints |
| `make test-all` | Run all tests |
| `make test-coverage` | Generate coverage report |

## Development

### Running Locally

```bash
go run cmd/main.go
```

### Running Tests

```bash
go test -v ./...
```

### Code Structure

- **Handler Layer:** Handles HTTP requests and responses
- **Repository Layer:** Manages data storage and retrieval
- **Model Layer:** Defines data structures

## Dependencies

- [gorilla/mux](https://github.com/gorilla/mux) - HTTP router
- [google/uuid](https://github.com/google/uuid) - UUID generation

## Notes

- The API uses in-memory storage, so data is lost when the server restarts
- The current implementation is not thread-safe for concurrent writes
- Passwords are stored in plain text (for demo purposes only)

## Future Improvements

- Add database persistence (PostgreSQL/MySQL)
- Implement proper authentication and authorization
- Add request validation and sanitization
- Implement thread-safe operations
- Add logging and monitoring
- Add API versioning
- Implement pagination for list operations

## License

This project is available under the MIT License.