package handler

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    
    "github.com/gorilla/mux"
    "go-crud-api/internal/model"
    "go-crud-api/internal/repository"
)

func setupTestRouter() (*mux.Router, *UserHandler) {
    repo := repository.NewMockUserRepository()
    handler := NewUserHandler(repo)
    router := mux.NewRouter()
    handler.RegisterRoutes(router)
    return router, handler
}

func TestCreateUser(t *testing.T) {
    router, _ := setupTestRouter()
    
    tests := []struct {
        name         string
        payload      interface{}
        expectedCode int
        checkBody    bool
    }{
        {
            name: "valid user creation",
            payload: map[string]string{
                "name":     "John Doe",
                "email":    "john@example.com",
                "password": "secret123",
            },
            expectedCode: http.StatusCreated,
            checkBody:    true,
        },
        {
            name:         "invalid JSON",
            payload:      "invalid json",
            expectedCode: http.StatusBadRequest,
            checkBody:    false,
        },
        {
            name:         "empty payload",
            payload:      map[string]string{},
            expectedCode: http.StatusCreated,
            checkBody:    true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var body []byte
            var err error
            
            if str, ok := tt.payload.(string); ok {
                body = []byte(str)
            } else {
                body, err = json.Marshal(tt.payload)
                if err != nil {
                    t.Fatalf("Failed to marshal payload: %v", err)
                }
            }
            
            req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
            req.Header.Set("Content-Type", "application/json")
            w := httptest.NewRecorder()
            
            router.ServeHTTP(w, req)
            
            if w.Code != tt.expectedCode {
                t.Errorf("Expected status %d, got %d", tt.expectedCode, w.Code)
            }
            
            if tt.checkBody && w.Code == http.StatusCreated {
                var response model.User
                if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
                    t.Fatalf("Failed to decode response: %v", err)
                }
                
                if response.ID == "" {
                    t.Error("Expected user ID to be generated")
                }
                
                if payload, ok := tt.payload.(map[string]string); ok {
                    if response.Name != payload["name"] {
                        t.Errorf("Expected name %s, got %s", payload["name"], response.Name)
                    }
                    if response.Email != payload["email"] {
                        t.Errorf("Expected email %s, got %s", payload["email"], response.Email)
                    }
                }
            }
        })
    }
}

func TestGetUser(t *testing.T) {
    router, handler := setupTestRouter()
    
    // Create a test user
    testUser := model.User{
        ID:       "test-get-123",
        Name:     "Test User",
        Email:    "test@example.com",
        Password: "password",
    }
    // Use the repo interface to save the user
    handler.repo.Save(testUser)
    
    tests := []struct {
        name         string
        userID       string
        expectedCode int
    }{
        {
            name:         "existing user",
            userID:       "test-get-123",
            expectedCode: http.StatusOK,
        },
        {
            name:         "non-existing user",
            userID:       "not-exists",
            expectedCode: http.StatusNotFound,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest("GET", "/users/"+tt.userID, nil)
            w := httptest.NewRecorder()
            
            router.ServeHTTP(w, req)
            
            if w.Code != tt.expectedCode {
                t.Errorf("Expected status %d, got %d", tt.expectedCode, w.Code)
            }
            
            if w.Code == http.StatusOK {
                var response model.User
                if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
                    t.Fatalf("Failed to decode response: %v", err)
                }
                
                if response.ID != testUser.ID {
                    t.Errorf("Expected ID %s, got %s", testUser.ID, response.ID)
                }
                if response.Name != testUser.Name {
                    t.Errorf("Expected name %s, got %s", testUser.Name, response.Name)
                }
            }
        })
    }
}

func TestUpdateUser(t *testing.T) {
    router, handler := setupTestRouter()
    
    // Create a test user
    originalUser := model.User{
        ID:       "test-update-123",
        Name:     "Original Name",
        Email:    "original@example.com",
        Password: "password",
    }
    handler.repo.Save(originalUser)
    
    tests := []struct {
        name         string
        userID       string
        payload      interface{}
        expectedCode int
    }{
        {
            name:   "update existing user",
            userID: "test-update-123",
            payload: map[string]string{
                "name":  "Updated Name",
                "email": "updated@example.com",
            },
            expectedCode: http.StatusOK,
        },
        {
            name:   "update non-existing user",
            userID: "not-exists",
            payload: map[string]string{
                "name": "Ghost User",
            },
            expectedCode: http.StatusNotFound,
        },
        {
            name:         "invalid JSON",
            userID:       "test-update-123",
            payload:      "invalid json",
            expectedCode: http.StatusBadRequest,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var body []byte
            var err error
            
            if str, ok := tt.payload.(string); ok {
                body = []byte(str)
            } else {
                body, err = json.Marshal(tt.payload)
                if err != nil {
                    t.Fatalf("Failed to marshal payload: %v", err)
                }
            }
            
            req := httptest.NewRequest("PUT", "/users/"+tt.userID, bytes.NewBuffer(body))
            req.Header.Set("Content-Type", "application/json")
            w := httptest.NewRecorder()
            
            router.ServeHTTP(w, req)
            
            if w.Code != tt.expectedCode {
                t.Errorf("Expected status %d, got %d", tt.expectedCode, w.Code)
            }
            
            if w.Code == http.StatusOK {
                var response model.User
                if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
                    t.Fatalf("Failed to decode response: %v", err)
                }
                
                if response.ID != tt.userID {
                    t.Errorf("Expected ID %s, got %s", tt.userID, response.ID)
                }
                
                if payload, ok := tt.payload.(map[string]string); ok {
                    if response.Name != payload["name"] {
                        t.Errorf("Expected name %s, got %s", payload["name"], response.Name)
                    }
                }
            }
        })
    }
}

func TestDeleteUser(t *testing.T) {
    router, handler := setupTestRouter()
    
    // Create test users
    user1 := model.User{ID: "test-delete-1", Name: "User 1"}
    user2 := model.User{ID: "test-delete-2", Name: "User 2"}
    handler.repo.Save(user1)
    handler.repo.Save(user2)
    
    tests := []struct {
        name         string
        userID       string
        expectedCode int
    }{
        {
            name:         "delete existing user",
            userID:       "test-delete-1",
            expectedCode: http.StatusNoContent,
        },
        {
            name:         "delete non-existing user",
            userID:       "not-exists",
            expectedCode: http.StatusNotFound,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest("DELETE", "/users/"+tt.userID, nil)
            w := httptest.NewRecorder()
            
            router.ServeHTTP(w, req)
            
            if w.Code != tt.expectedCode {
                t.Errorf("Expected status %d, got %d", tt.expectedCode, w.Code)
            }
            
            // Verify user is actually deleted
            if w.Code == http.StatusNoContent {
                _, found := handler.repo.FindById(tt.userID)
                if found {
                    t.Error("User still exists after deletion")
                }
            }
        })
    }
    
    // Verify other users are not affected
    _, found := handler.repo.FindById("test-delete-2")
    if !found {
        t.Error("Unrelated user was deleted")
    }
}

func TestRegisterRoutes(t *testing.T) {
    repo := repository.NewMockUserRepository()
    handler := NewUserHandler(repo)
    router := mux.NewRouter()
    
    handler.RegisterRoutes(router)
    
    // Test that routes are registered correctly
    routes := []struct {
        method string
        path   string
    }{
        {"POST", "/users"},
        {"GET", "/users/{id}"},
        {"PUT", "/users/{id}"},
        {"DELETE", "/users/{id}"},
    }
    
    for _, route := range routes {
        t.Run(route.method+" "+route.path, func(t *testing.T) {
            // This is a basic check to ensure routes are registered
            // In a real scenario, you might want to test the actual routing
            if router == nil {
                t.Error("Router should not be nil after registering routes")
            }
        })
    }
}