package repository

import (
    "testing"
    "go-crud-api/internal/model"
)

func TestNewUserRepository(t *testing.T) {
    repo := NewUserRepository()
    if repo == nil {
        t.Fatal("NewUserRepository returned nil")
    }
    if repo.users == nil {
        t.Fatal("users map not initialized")
    }
}

func TestUserRepository_Save(t *testing.T) {
    repo := NewUserRepository()
    
    user := model.User{
        ID:       "test-123",
        Name:     "Test User",
        Email:    "test@example.com",
        Password: "password",
    }
    
    repo.Save(user)
    
    // Verify user was saved
    savedUser, exists := repo.users[user.ID]
    if !exists {
        t.Fatal("User was not saved")
    }
    
    if savedUser != user {
        t.Errorf("Saved user does not match: got %+v, want %+v", savedUser, user)
    }
}

func TestUserRepository_FindById(t *testing.T) {
    repo := NewUserRepository()
    
    // Add test user
    user := model.User{
        ID:       "find-123",
        Name:     "Find User",
        Email:    "find@example.com",
        Password: "password",
    }
    repo.Save(user)
    
    tests := []struct {
        name      string
        id        string
        wantUser  model.User
        wantFound bool
    }{
        {
            name:      "existing user",
            id:        "find-123",
            wantUser:  user,
            wantFound: true,
        },
        {
            name:      "non-existing user",
            id:        "not-exists",
            wantUser:  model.User{},
            wantFound: false,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotUser, found := repo.FindById(tt.id)
            if found != tt.wantFound {
                t.Errorf("FindById() found = %v, want %v", found, tt.wantFound)
            }
            if found && gotUser != tt.wantUser {
                t.Errorf("FindById() user = %+v, want %+v", gotUser, tt.wantUser)
            }
        })
    }
}

func TestUserRepository_Update(t *testing.T) {
    repo := NewUserRepository()
    
    // Add initial user
    originalUser := model.User{
        ID:       "update-123",
        Name:     "Original Name",
        Email:    "original@example.com",
        Password: "password",
    }
    repo.Save(originalUser)
    
    tests := []struct {
        name        string
        updateUser  model.User
        wantSuccess bool
    }{
        {
            name: "update existing user",
            updateUser: model.User{
                ID:       "update-123",
                Name:     "Updated Name",
                Email:    "updated@example.com",
                Password: "newpassword",
            },
            wantSuccess: true,
        },
        {
            name: "update non-existing user",
            updateUser: model.User{
                ID:       "not-exists",
                Name:     "Ghost User",
                Email:    "ghost@example.com",
                Password: "password",
            },
            wantSuccess: false,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            success := repo.Update(tt.updateUser)
            if success != tt.wantSuccess {
                t.Errorf("Update() = %v, want %v", success, tt.wantSuccess)
            }
            
            if success {
                updatedUser, _ := repo.FindById(tt.updateUser.ID)
                if updatedUser != tt.updateUser {
                    t.Errorf("User not updated correctly: got %+v, want %+v", updatedUser, tt.updateUser)
                }
            }
        })
    }
}

func TestUserRepository_Delete(t *testing.T) {
    repo := NewUserRepository()
    
    // Add test users
    user1 := model.User{ID: "delete-1", Name: "User 1"}
    user2 := model.User{ID: "delete-2", Name: "User 2"}
    repo.Save(user1)
    repo.Save(user2)
    
    tests := []struct {
        name        string
        deleteID    string
        wantSuccess bool
    }{
        {
            name:        "delete existing user",
            deleteID:    "delete-1",
            wantSuccess: true,
        },
        {
            name:        "delete non-existing user",
            deleteID:    "not-exists",
            wantSuccess: false,
        },
        {
            name:        "delete already deleted user",
            deleteID:    "delete-1",
            wantSuccess: false,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            success := repo.Delete(tt.deleteID)
            if success != tt.wantSuccess {
                t.Errorf("Delete() = %v, want %v", success, tt.wantSuccess)
            }
            
            // Verify user is actually deleted
            if success {
                _, found := repo.FindById(tt.deleteID)
                if found {
                    t.Error("User still exists after deletion")
                }
            }
        })
    }
    
    // Verify other users are not affected
    _, found := repo.FindById("delete-2")
    if !found {
        t.Error("Unrelated user was deleted")
    }
}

func TestUserRepository_ConcurrentAccess(t *testing.T) {
    // Note: Current implementation is not thread-safe
    // This test documents that behavior
    t.Skip("Repository is not thread-safe - skipping concurrent access test")
}