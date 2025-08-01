package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "go-crud-api/internal/database"
    "go-crud-api/internal/handler"
    "go-crud-api/internal/middleware"
    "go-crud-api/internal/repository"
)

func main() {
    // Initialize database connection
    db, err := database.NewMySQLConnection()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    r := mux.NewRouter()

    userRepo := repository.NewUserRepository(db)
    userHandler := handler.NewUserHandler(userRepo)

    r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
    r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
    r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

    // Apply CORS middleware
    handler := middleware.CORS(r)
    
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", handler); err != nil {
        log.Fatal(err)
    }
}