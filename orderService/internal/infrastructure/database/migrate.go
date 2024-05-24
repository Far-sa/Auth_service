package database

import (
    "database/sql"
    "log"
)

func Migrate(db *sql.DB) error {
    migration := `
    CREATE TABLE IF NOT EXISTS orders (
        order_id VARCHAR(255) NOT NULL,
        user_id VARCHAR(255) NOT NULL,
        status VARCHAR(255) NOT NULL,
        total_amount NUMERIC(10, 2) NOT NULL,
        created_at TIMESTAMP NOT NULL,
        PRIMARY KEY (order_id)
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
