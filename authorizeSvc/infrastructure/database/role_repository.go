package database

import (
	"authorization-service/internal/interfaces"
	"context"
	"log"
)

type PostgresRoleRepository struct {
	db *sqlDB
}

func NewPostgresRoleRepository(db *sqlDB) interfaces.RoleRepository {
	return &PostgresRoleRepository{db}
}

func (r *PostgresRoleRepository) AssignRole(ctx context.Context, username, role string) error {
	query := `INSERT INTO roles (username, role) VALUES ($1, $2) ON CONFLICT (username) DO UPDATE SET role = $2`
	_, err := r.db.ExecContext(ctx, query, username, role)
	if err != nil {
		log.Printf("Error assigning role: %v", err)
		return err
	}
	return nil
}

func (r *PostgresRoleRepository) CheckPermission(ctx context.Context, username, permission string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM permissions WHERE username = $1 AND permission = $2)`
	var hasPermission bool
	err := r.db.QueryRowContext(ctx, query, username, permission).Scan(&hasPermission)
	if err != nil {
		log.Printf("Error checking permission: %v", err)
		return false, err
	}
	return hasPermission, nil
}

func (r *PostgresRoleRepository) UpdateUserRoles(ctx context.Context, userID string, role string) error {
	log.Printf("Updating roles for user: %s to role: %s", userID, role)
	// Logic to update roles in the database
	return nil
}
