package repository

import (
    "go-crud-api/internal/model"
)

type UserRepository struct {
    users map[string]model.User
}

func NewUserRepository() *UserRepository {
    return &UserRepository{
        users: make(map[string]model.User),
    }
}

func (r *UserRepository) Save(user model.User) {
    r.users[user.ID] = user
}

func (r *UserRepository) FindById(id string) (model.User, bool) {
    user, exists := r.users[id]
    return user, exists
}

func (r *UserRepository) Update(user model.User) bool {
    _, exists := r.users[user.ID]
    if exists {
        r.users[user.ID] = user
    }
    return exists
}

func (r *UserRepository) Delete(id string) bool {
    _, exists := r.users[id]
    if exists {
        delete(r.users, id)
    }
    return exists
}