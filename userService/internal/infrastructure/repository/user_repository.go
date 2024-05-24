package database

import (
	"database/sql"
	"user-service/interfaces"
	"user-service/internal/domain/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) interfaces.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user models.User) error {
	query := `INSERT INTO users (user_id, name, email, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, user.UserID, user.Name, user.Email, user.CreatedAt)
	return err
}

func (r *UserRepository) GetUser(userID string) (models.User, error) {
	var user models.User
	query := `SELECT user_id, name, email, created_at FROM users WHERE user_id = $1`
	row := r.db.QueryRow(query, userID)
	err := row.Scan(&user.UserID, &user.Name, &user.Email, &user.CreatedAt)
	return user, err
}
