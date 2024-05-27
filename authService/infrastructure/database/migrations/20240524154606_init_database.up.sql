CREATE TABLE tokens (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL, -- This is a reference to a user ID in the User Service
    token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL
);