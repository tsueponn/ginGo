CREATE TYPE IF NOT EXISTS user_role AS ENUM ('admin', 'user');

CREATE TABLE IF NOT EXISTS users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(100) UNIQUE NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       role user_role DEFAULT 'user',
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP
);