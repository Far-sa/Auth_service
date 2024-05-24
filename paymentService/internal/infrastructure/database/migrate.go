package database

import (
	"database/sql"
)

func Migrate(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS payments (
        payment_id VARCHAR(50) PRIMARY KEY,
        order_id VARCHAR(50) NOT NULL,
        amount FLOAT NOT NULL,
        method VARCHAR(50) NOT NULL,
        created_at TIMESTAMP NOT NULL
    );`
	_, err := db.Exec(query)
	return err
}
