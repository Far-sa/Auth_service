CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT
);

CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT
);

CREATE TABLE user_roles (
    user_id INT NOT NULL, -- This is a reference to a user ID in the User Service
    role_id INT NOT NULL, -- This is a reference to a role ID in the Authorize Service
    PRIMARY KEY (user_id, role_id)
);

CREATE TABLE role_permissions (
    role_id INT NOT NULL, -- This is a reference to a role ID in the Authorize Service
    permission_id INT NOT NULL, -- This is a reference to a permission ID in the Authorize Service
    PRIMARY KEY (role_id, permission_id)
);
