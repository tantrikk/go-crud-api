package handler

import (
    "net/http"
    "github.com/gorilla/mux"
    "go-crud-api/internal/model"
    "go-crud-api/internal/repository"
)

type UserHandler struct {
    repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
    return &UserHandler{repo: repo}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    // Implementation for creating a user
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    // Implementation for getting a user
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    // Implementation for updating a user
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    // Implementation for deleting a user
}

func (h *UserHandler) RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/users", h.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
    r.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")
}