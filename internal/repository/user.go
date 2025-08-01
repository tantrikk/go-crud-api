package repository

import (
    "database/sql"
    "go-crud-api/internal/model"
    "go-crud-api/internal/database"
)

type UserRepository struct {
    db *database.MySQLDB
}

func NewUserRepository(db *database.MySQLDB) *UserRepository {
    return &UserRepository{
        db: db,
    }
}

func (r *UserRepository) GetAll() ([]model.User, error) {
    query := `SELECT id, name, email, password FROM users`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []model.User
    for rows.Next() {
        var user model.User
        err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    
    return users, nil
}

func (r *UserRepository) Save(user model.User) error {
    query := `INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)`
    _, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Password)
    return err
}

func (r *UserRepository) FindById(id string) (model.User, bool) {
    var user model.User
    query := `SELECT id, name, email, password FROM users WHERE id = ?`
    err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
    
    if err != nil {
        if err == sql.ErrNoRows {
            return user, false
        }
        return user, false
    }
    
    return user, true
}

func (r *UserRepository) Update(user model.User) bool {
    query := `UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?`
    result, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.ID)
    
    if err != nil {
        return false
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return false
    }
    
    return rowsAffected > 0
}

func (r *UserRepository) Delete(id string) bool {
    query := `DELETE FROM users WHERE id = ?`
    result, err := r.db.Exec(query, id)
    
    if err != nil {
        return false
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return false
    }
    
    return rowsAffected > 0
}