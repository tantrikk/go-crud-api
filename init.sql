CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email)
);

-- Insert some sample data (optional)
INSERT INTO users (id, name, email, password) VALUES 
    ('550e8400-e29b-41d4-a716-446655440001', 'Admin User', 'admin@example.com', 'admin123'),
    ('550e8400-e29b-41d4-a716-446655440002', 'Test User', 'test@example.com', 'test123')
ON DUPLICATE KEY UPDATE name=name;