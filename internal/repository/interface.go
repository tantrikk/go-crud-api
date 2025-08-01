package repository

import "go-crud-api/internal/model"

// UserRepositoryInterface defines the methods for user repository
type UserRepositoryInterface interface {
    GetAll() ([]model.User, error)
    Save(user model.User) error
    FindById(id string) (model.User, bool)
    Update(user model.User) bool
    Delete(id string) bool
}