package database

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
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

func Migrate(db *sql.DB) {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    )`
    if _, err := db.Exec(query); err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
}
