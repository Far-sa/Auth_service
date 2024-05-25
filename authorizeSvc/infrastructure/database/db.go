package database

import (
	"context"
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

type sqlDB struct {
	db *sql.DB
}

func NewSQLDB(dataSourceName string) (*sqlDB, error) {
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

	return &sqlDB{db: db}, nil
}

func (db *sqlDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	stmt, err := db.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close() // Close prepared statement after execution

	result, err := stmt.ExecContext(ctx, args...)
	return result, err
}

func (db *sqlDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	stmt, err := db.db.PrepareContext(ctx, query)
	if err != nil {
		return nil
	}
	defer stmt.Close() // Close prepared statement after execution

	return stmt.QueryRowContext(ctx, args...)
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
