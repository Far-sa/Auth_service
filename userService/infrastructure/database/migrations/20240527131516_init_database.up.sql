CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password CHAR(60) NOT NULL,
    full_name VARCHAR(100),
    phone_number VARCHAR(15),
    birthdate DATE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);