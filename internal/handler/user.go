package handler

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "go-crud-api/internal/model"
    "go-crud-api/internal/repository"
    "github.com/google/uuid"
)

type UserHandler struct {
    repo repository.UserRepositoryInterface
}

func NewUserHandler(repo repository.UserRepositoryInterface) *UserHandler {
    return &UserHandler{repo: repo}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user model.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    
    user.ID = uuid.New().String()
    if err := h.repo.Save(user); err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := h.repo.GetAll()
    if err != nil {
        http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    user, exists := h.repo.FindById(id)
    if !exists {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    var user model.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    
    user.ID = id
    if !h.repo.Update(user) {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    if !h.repo.Delete(id) {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/users", h.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
    r.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")
}