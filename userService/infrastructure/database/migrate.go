package database

import (
	"database/sql"
)

func Migrate(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        user_id VARCHAR(50) PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL,
        created_at TIMESTAMP NOT NULL
    );`
	_, err := db.Exec(query)
	return err
}
