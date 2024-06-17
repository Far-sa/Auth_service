package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"user-service/internal/entity"

	"github.com/lib/pq"
)

type SqlDB struct {
	db *sql.DB
}

// Add this method to the sqlDB type.
func (d *SqlDB) Conn() *sql.DB {
	return d.db
}

func NewSQLDB(dsn string) (*SqlDB, error) {

	//dsn := "postgres://root:password@postgres-auth:5432/auth-db?sslmode=disable" // Connect to the database directly
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Errorf("can not open postgres database: %v", err))
	}

	// Set connection pool parameters (optional)
	db.SetMaxOpenConns(20) // Adjust as needed
	db.SetMaxIdleConns(10) // Adjust as needed

	// Perform initial ping to check connectivity
	if err := db.PingContext(context.Background()); err != nil {
		return nil, errors.New("failed to ping database: " + err.Error())
	}

	return &SqlDB{db: db}, nil
}

func (db *SqlDB) GetUserByID(ctx context.Context, userID string) (*entity.UserProfile, error) {
	var user entity.UserProfile
	query := `SELECT id, full_name, username, email, created_at, birthdate FROM users WHERE user_id = $1`
	row := db.db.QueryRowContext(ctx, query, userID)
	err := row.Scan(&user.ID, &user.FullName, &user.Username, &user.Email, &user.CreatedAt, &user.Birthdate)
	if err != nil {
		if _, ok := err.(*pq.Error); ok { // Handle specific postgres errors if needed
			// Handle specific postgres error here
			return nil, fmt.Errorf("error getting user: %w", err)
		}
		return nil, fmt.Errorf("error scanning user: %w", err)
	}
	return &user, nil
}
