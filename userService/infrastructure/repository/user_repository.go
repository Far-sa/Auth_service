package repository

import (
	"context"
	"user-service/internal/entity"
)

func (r *DB) GetUserByID(ctx context.Context, userID string) (*entity.UserProfile, error) {
	var user entity.UserProfile
	query := `SELECT id, full_name,username, email, created_at ,birthdate FROM users WHERE user_id = $1`
	row := r.conn.Conn().QueryRowContext(ctx, query, userID)
	err := row.Scan(&user.ID, &user.FullName, &user.Username, &user.Email, user.Birthdate, &user.CreatedAt)
	return &user, err
}

func (r *DB) FindUserByEmail(ctx context.Context, Email string) (*entity.UserProfile, error) {
	var user entity.UserProfile
	query := `SELECT id, full_name,username, email, created_at ,birthdate FROM users WHERE username = $1 OR email = $1`
	row := r.conn.Conn().QueryRowContext(ctx, query, Email)
	err := row.Scan(&user.ID, &user.FullName, &user.Username, &user.Email, user.Birthdate, &user.CreatedAt)
	return &user, err
}

func (r *DB) CreateUser(ctx context.Context, user *entity.UserProfile) (*entity.UserProfile, error) {
	query := `INSERT INTO users (full_name, username, email, password, phone_number, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	row := r.conn.Conn().QueryRowContext(ctx, query, user.FullName, user.Username, user.Email, user.Password, user.PhoneNumber, user.CreatedAt)
	err := row.Scan(&user.ID)
	return user, err
}
