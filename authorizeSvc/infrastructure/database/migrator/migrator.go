package migrator

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// MigrateInterface is an interface that matches the migrate.Migrate methods used.
type MigrateInterface interface {
	Up() error
	Down() error
}

type Migrator struct {
	m *migrate.Migrate
}

func NewMigrator(db *sql.DB, migrationFilesPath string) (*Migrator, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{}) // Convert db to *sql.DB
	if err != nil {
		return nil, fmt.Errorf("could not create database driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationFilesPath),
		"postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("migration failed to initialize: %w", err)
	}

	return &Migrator{m: m}, nil
}

func (m *Migrator) Up() error {
	err := m.m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("an error occurred while migrating up: %w", err)
	}
	return nil
}

func (m *Migrator) Down() error {
	err := m.m.Down()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("an error occurred while migrating down: %w", err)
	}
	return nil
}
