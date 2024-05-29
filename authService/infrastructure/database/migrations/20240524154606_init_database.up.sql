
-- Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- Tokens table
CREATE TABLE tokens (
    id SERIAL PRIMARY KEY,
    access_token TEXT NOT NULL,
    refresh_token TEXT NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id),
    access_token_expires_at TIMESTAMP NOT NULL,
    refresh_token_expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- Tokens table
-- CREATE TABLE tokens (
--     id SERIAL PRIMARY KEY,
--     token TEXT NOT NULL,
--     user_id INTEGER NOT NULL REFERENCES users(id),
--     expires_at TIMESTAMP NOT NULL,
--     created_at TIMESTAMP NOT NULL
-- );

-- CREATE TABLE tokens (
--     id SERIAL PRIMARY KEY,
--     user_id INT NOT NULL, -- This is a reference to a user ID in the User Service
--     token VARCHAR(255) NOT NULL,
--     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     expires_at TIMESTAMP NOT NULL
-- );

