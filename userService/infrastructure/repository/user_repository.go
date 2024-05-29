package repository

import (
	"context"
	"user-service/internal/entity"
)

func (r *DB) GetUserByID(userID string) (entity.UserProfile, error) {
	var user entity.UserProfile
	query := `SELECT id, full_name,username, email, created_at ,birthdate FROM users WHERE user_id = $1`
	row := r.conn.Conn().QueryRow(query, userID)
	err := row.Scan(&user.ID, &user.FullName, &user.Username, &user.Email, user.Birthdate, &user.CreatedAt)
	return user, err
}

func (r *DB) FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*entity.UserProfile, error) {
	var user entity.UserProfile
	query := `SELECT id, full_name,username, email, created_at ,birthdate FROM users WHERE username = $1 OR email = $1`
	row := r.conn.Conn().QueryRowContext(ctx, query, usernameOrEmail)
	err := row.Scan(&user.ID, &user.FullName, &user.Username, &user.Email, user.Birthdate, &user.CreatedAt)
	return &user, err
}
