package database

import (
	"context"
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

type SqlDB struct {
	db *sql.DB
}

// Add this method to the sqlDB type.
func (s *SqlDB) Conn() *sql.DB {
	return s.db
}

func NewSQLDB(dataSourceName string) (*SqlDB, error) {
	db, err := sql.Open("postgres", dataSourceName) // Adjust driver as needed
	if err != nil {
		return nil, err
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

const createTableQueries = `
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    permission VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

func Migrate(db *sql.DB) error {
	_, err := db.Exec(createTableQueries)
	if err != nil {
		return err
	}
	log.Println("Database migration completed successfully.")
	return nil
}
