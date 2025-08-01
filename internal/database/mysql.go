package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
    *sql.DB
}

func NewMySQLConnection() (*MySQLDB, error) {
    dbHost := getEnv("DB_HOST", "localhost")
    dbPort := getEnv("DB_PORT", "3306")
    dbUser := getEnv("DB_USER", "apiuser")
    dbPassword := getEnv("DB_PASSWORD", "apipassword")
    dbName := getEnv("DB_NAME", "userdb")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", 
        dbUser, dbPassword, dbHost, dbPort, dbName)

    var db *sql.DB
    var err error

    // Retry connection to handle container startup delays
    for i := 0; i < 30; i++ {
        db, err = sql.Open("mysql", dsn)
        if err != nil {
            log.Printf("Failed to open database: %v", err)
            time.Sleep(1 * time.Second)
            continue
        }

        err = db.Ping()
        if err == nil {
            break
        }

        log.Printf("Failed to ping database (attempt %d/30): %v", i+1, err)
        time.Sleep(1 * time.Second)
    }

    if err != nil {
        return nil, fmt.Errorf("failed to connect to database after 30 attempts: %v", err)
    }

    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)

    log.Println("Successfully connected to MySQL database")

    return &MySQLDB{db}, nil
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}