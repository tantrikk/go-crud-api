package model

import (
    "encoding/json"
    "testing"
)

func TestUserJSONMarshaling(t *testing.T) {
    tests := []struct {
        name string
        user User
        want string
    }{
        {
            name: "complete user",
            user: User{
                ID:       "123",
                Name:     "John Doe",
                Email:    "john@example.com",
                Password: "secret",
            },
            want: `{"id":"123","name":"John Doe","email":"john@example.com","password":"secret"}`,
        },
        {
            name: "user without password in response",
            user: User{
                ID:       "456",
                Name:     "Jane Doe",
                Email:    "jane@example.com",
                Password: "",
            },
            want: `{"id":"456","name":"Jane Doe","email":"jane@example.com"}`,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := json.Marshal(tt.user)
            if err != nil {
                t.Fatalf("Failed to marshal user: %v", err)
            }
            if string(got) != tt.want {
                t.Errorf("Got %s, want %s", string(got), tt.want)
            }
        })
    }
}

func TestUserJSONUnmarshaling(t *testing.T) {
    tests := []struct {
        name    string
        json    string
        want    User
        wantErr bool
    }{
        {
            name: "valid user JSON",
            json: `{"id":"789","name":"Bob Smith","email":"bob@example.com","password":"pass123"}`,
            want: User{
                ID:       "789",
                Name:     "Bob Smith",
                Email:    "bob@example.com",
                Password: "pass123",
            },
            wantErr: false,
        },
        {
            name: "partial user JSON",
            json: `{"name":"Alice","email":"alice@example.com"}`,
            want: User{
                Name:  "Alice",
                Email: "alice@example.com",
            },
            wantErr: false,
        },
        {
            name:    "invalid JSON",
            json:    `{"name":}`,
            want:    User{},
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var got User
            err := json.Unmarshal([]byte(tt.json), &got)
            if (err != nil) != tt.wantErr {
                t.Errorf("Unmarshal error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !tt.wantErr && got != tt.want {
                t.Errorf("Got %+v, want %+v", got, tt.want)
            }
        })
    }
}