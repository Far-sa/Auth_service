package repository

import (
	"authentication-service/domain/entities"
	"context"
	"database/sql"
)

// type PostgresUserRepository struct {
// 	db *SqlDB
// }

// func NewPostgresUserRepository(db *SqlDB) interfaces.UserRepository {
// 	return &PostgresUserRepository{db: db}
// }

func (r *DB) Save(ctx context.Context, user *entities.User) error {
	// Prepared statement for saving a user
	query := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)`

	// Execute the prepared statement with user data
	_, err := r.conn.Conn().ExecContext(ctx, query, user.Username, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}

func (r *DB) FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*entities.User, error) {
	// Prepared statement for finding a user by username or email
	query := `SELECT * FROM users WHERE username = $1 OR email = $2`

	// Execute the prepared statement with username/email
	row := r.conn.Conn().QueryRowContext(ctx, query, usernameOrEmail, usernameOrEmail)

	// Scan the row into a user object
	var u entities.User
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &u, nil
}

// ... Implementations for other user repository methods using r.db object and reusable functions
