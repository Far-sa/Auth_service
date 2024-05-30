package repository

import (
	"authentication-service/domain/entities"
	"context"
)

// type PostgresUserRepository struct {
// 	db *SqlDB
// }

// func NewPostgresUserRepository(db *SqlDB) interfaces.UserRepository {
// 	return &PostgresUserRepository{db: db}
// }

func (r *DB) SaveToken(ctx context.Context, token *entities.TokenPair) error {
	query := `INSERT INTO tokens (user_id, access_token, refresh_token) VALUES ($1, $2, $3)`
	_, err := r.conn.Conn().ExecContext(ctx, query, token.AccessToken, token.RefreshToken)
	return err

}

func (r *DB) FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*entities.User, error) {
	return nil, nil
}

func (r *DB) SaveUser(ctx context.Context, user entities.User) error {

	query := "INSERT INTO users (username, email,password) VALUES($1,$2,$3)"

	_, err := r.conn.Conn().ExecContext(ctx, query, user.Username, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}
	return nil
}

// func (r *DB) FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*entities.User, error) {
// 	// Prepared statement for finding a user by username or email
// 	query := `SELECT * FROM users WHERE username = $1 OR email = $2`

// 	// Execute the prepared statement with username/email
// 	row := r.conn.Conn().QueryRowContext(ctx, query, usernameOrEmail, usernameOrEmail)

// 	// Scan the row into a user object
// 	var u entities.User
// 	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil // User not found
// 		}
// 		return nil, err
// 	}

// 	return &u, nil
// }

// ... Implementations for other user repository methods using r.db object and reusable functions
