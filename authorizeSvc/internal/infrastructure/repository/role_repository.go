package repository

import (
	"authorization-service/interfaces"
	"database/sql"
	"log"
)

type RoleRepository interface {
	AssignRole(username, role string) error
	UpdateUserRoles(userID string, role string) error
	CheckPermission(username, permission string) (bool, error)
}

type PostgresRoleRepository struct {
	db *sql.DB
}

func NewPostgresRoleRepository(db *sql.DB) interfaces.RoleRepository {
	return &PostgresRoleRepository{db}
}

func (r *PostgresRoleRepository) AssignRole(username, role string) error {
	query := `INSERT INTO roles (username, role) VALUES ($1, $2) ON CONFLICT (username) DO UPDATE SET role = $2`
	_, err := r.db.Exec(query, username, role)
	if err != nil {
		log.Printf("Error assigning role: %v", err)
		return err
	}
	return nil
}

func (r *PostgresRoleRepository) CheckPermission(username, permission string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM permissions WHERE username = $1 AND permission = $2)`
	var hasPermission bool
	err := r.db.QueryRow(query, username, permission).Scan(&hasPermission)
	if err != nil {
		log.Printf("Error checking permission: %v", err)
		return false, err
	}
	return hasPermission, nil
}

func (r *PostgresRoleRepository) UpdateUserRoles(userID string, role string) error {
	log.Printf("Updating roles for user: %s to role: %s", userID, role)
	// Logic to update roles in the database
	return nil
}
