package repository

import (
	"user-service/internal/entity"
)

func (r *DB) CreateUser(user entity.User) error {
	query := `INSERT INTO users (user_id, name, email, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.conn.Conn().Exec(query, user.UserID, user.Name, user.Email, user.CreatedAt)
	return err
}

func (r *DB) GetUser(userID string) (entity.User, error) {
	var user entity.User
	query := `SELECT user_id, name, email, created_at FROM users WHERE user_id = $1`
	row := r.conn.Conn().QueryRow(query, userID)
	err := row.Scan(&user.UserID, &user.Name, &user.Email, &user.CreatedAt)
	return user, err
}
