package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "go-crud-api/internal/handler"
    "go-crud-api/internal/repository"
)

func main() {
    r := mux.NewRouter()

    userRepo := repository.NewUserRepository()
    userHandler := handler.NewUserHandler(userRepo)

    r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
    r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}