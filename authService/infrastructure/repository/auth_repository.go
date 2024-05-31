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

func (r *DB) SaveToken(ctx context.Context, token *entities.Token) error {
	query := `INSERT INTO tokens (id, user_id, access_token, refresh_token,created_at,access_token_expires_at,refresh_token_expires_at) VALUES ($1, $2, $3,$4,$5,$6,$7)`
	_, err := r.conn.Conn().ExecContext(ctx, query, token.AccessToken, token.RefreshToken, token.CreatedAt,
		token.AccessTokenExpiresAt, token.RefreshTokenExpiresAt, token.UserID)
	return err

}

//! move to user service
// func (r *DB) FindByUserEmail(ctx context.Context, Email string) (*entities.User, error) {
// 	query := `SELECT * FROM users WHERE email = $1`
// 	row := r.conn.Conn().QueryRowContext(ctx, query, Email)
// 	user := &entities.User{}
// 	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func (r *DB) SaveUser(ctx context.Context, user *entities.User) error {

// 	query := "INSERT INTO users (username, email,password) VALUES($1,$2,$3)"

// 	_, err := r.conn.Conn().ExecContext(ctx, query, user.Username, user.Email, user.PasswordHash)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
