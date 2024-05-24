package database

import (
	"database/sql"
)

func Migrate(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS inventory (
        product_id VARCHAR(50) PRIMARY KEY,
        quantity INT NOT NULL,
        last_updated TIMESTAMP NOT NULL
    );`
	_, err := db.Exec(query)
	return err
}
