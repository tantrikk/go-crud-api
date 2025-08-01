<<<<<<< HEAD
# go-crud-api
go crud api
=======
# Go CRUD API

This project is a simple CRUD API for managing users, built with Go. It provides basic operations to create, read, update, and delete user information.

## Project Structure

```
go-crud-api
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── handler
│   │   └── user.go      # HTTP request handlers for user operations
│   ├── model
│   │   └── user.go      # User model definition
│   └── repository
│       └── user.go      # Data storage interaction for users
├── go.mod                # Module definition
├── go.sum                # Dependency checksums
├── Dockerfile            # Docker image instructions
└── README.md             # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd go-crud-api
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/main.go
   ```

## Usage

The API provides the following endpoints:

- **Create User**
  - `POST /users`
  - Request Body: `{ "name": "John Doe", "email": "john@example.com", "password": "securepassword" }`

- **Get User**
  - `GET /users/{id}`
  
- **Update User**
  - `PUT /users/{id}`
  - Request Body: `{ "name": "John Doe Updated", "email": "john.updated@example.com" }`

- **Delete User**
  - `DELETE /users/{id}`

## Docker

To build and run the application in a Docker container, use the following commands:

1. **Build the Docker image:**
   ```
   docker build -t go-crud-api .
   ```

2. **Run the Docker container:**
   ```
   docker run -p 8080:8080 go-crud-api
   ```

The application will be accessible at `http://localhost:8080`.
>>>>>>> f05c9c1 (initial commit)
