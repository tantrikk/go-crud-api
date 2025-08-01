package repository

import "go-crud-api/internal/model"

// MockUserRepository implements an in-memory version for testing
type MockUserRepository struct {
    users map[string]model.User
}

func NewMockUserRepository() *MockUserRepository {
    return &MockUserRepository{
        users: make(map[string]model.User),
    }
}

func (r *MockUserRepository) GetAll() ([]model.User, error) {
    users := make([]model.User, 0, len(r.users))
    for _, user := range r.users {
        users = append(users, user)
    }
    return users, nil
}

func (r *MockUserRepository) Save(user model.User) error {
    r.users[user.ID] = user
    return nil
}

func (r *MockUserRepository) FindById(id string) (model.User, bool) {
    user, exists := r.users[id]
    return user, exists
}

func (r *MockUserRepository) Update(user model.User) bool {
    _, exists := r.users[user.ID]
    if exists {
        r.users[user.ID] = user
    }
    return exists
}

func (r *MockUserRepository) Delete(id string) bool {
    _, exists := r.users[id]
    if exists {
        delete(r.users, id)
    }
    return exists
}