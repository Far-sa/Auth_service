package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *sql.DB) error {
	migration := `
    CREATE TABLE IF NOT EXISTS cart (
        user_id VARCHAR(255) NOT NULL,
        product_id VARCHAR(255) NOT NULL,
        quantity INT NOT NULL,
        PRIMARY KEY (user_id, product_id)
    );
    `

	_, err := db.Exec(migration)
	if err != nil {
		log.Fatalf("failed to execute migration: %v", err)
		return err
	}

	log.Println("Database migrated successfully.")
	return nil
}
