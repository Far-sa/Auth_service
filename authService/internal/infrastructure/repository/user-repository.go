package repository

import (
	"authentication-service/interfaces"
	"authentication-service/internal/domain"
	"database/sql"
	"errors"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) interfaces.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Save(user *domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, username, password) VALUES ($1, $2, $3)", user.ID, user.Username, user.Password)
	if err != nil {
		return ErrUserAlreadyExists
	}
	return nil
}

func (r *PostgresUserRepository) FindByUsername(username string) (*domain.User, error) {
	row := r.db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return user, nil
}
