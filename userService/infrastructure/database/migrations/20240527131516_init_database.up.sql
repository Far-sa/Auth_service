
-- CREATE TABLE users (
--     id SERIAL PRIMARY KEY,
--     username VARCHAR(255) UNIQUE NOT NULL,
--     email VARCHAR(255) UNIQUE NOT NULL,
--     password_hash VARCHAR(255) NOT NULL,
--     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
-- );

-- UserProfiles table
CREATE TABLE user_profiles (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    full_name VARCHAR(255),
    birthdate DATE,
    created_at TIMESTAMP NOT NULL
);