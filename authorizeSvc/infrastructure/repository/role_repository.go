package repository

import (
	"authorization-service/internal/entity"
	"authorization-service/internal/interfaces"
	"context"
	"errors"
	"fmt"
	"log"
)

// type PostgresRoleRepository struct {
// 	db *sqlDB
// }

// func NewPostgresRoleRepository(db *sqlDB) interfaces.RoleRepository {
// 	return &PostgresRoleRepository{db}
// }

func (r *DB) AssignRole(ctx context.Context, userID, role string) error {
	query := `INSERT INTO roles (user_id, role) VALUES ($1, $2) ON CONFLICT (username) DO UPDATE SET role = $2`
	_, err := r.conn.Conn().ExecContext(ctx, query, userID, role)
	if err != nil {
		log.Printf("Error assigning role: %v", err)
		return err
	}
	return nil
}

func (r *DB) CheckPermission(ctx context.Context, username, permission string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM permissions WHERE username = $1 AND permission = $2)`
	var hasPermission bool
	err := r.conn.Conn().QueryRowContext(ctx, query, username, permission).Scan(&hasPermission)
	if err != nil {
		log.Printf("Error checking permission: %v", err)
		return false, err
	}
	return hasPermission, nil
}

func (r *DB) UpdateUserRoles(ctx context.Context, userID string, role string) error {
	log.Printf("Updating roles for user: %s to role: %s", userID, role)
	query := "update users set role = $1 where id = $2"
	row := r.conn.Conn().QueryRowContext(ctx, query, role, userID)
	err := row.Err()

	if err != nil {
		// Log and return the error if the query execution failed
		err = fmt.Errorf("error updating roles for user %s: %w", userID, err)
		return err
	}
	return nil
}

func (r *DB) GetRoleByUserID(ctx context.Context, userID string) (entity.Role, error) {
	query := "SELECT * FROM roles WHERE user_id = $1"
	row := r.conn.Conn().QueryRowContext(ctx, query, userID)
	err := row.Err() // Add this line to declare and assign the error variable
	if err != nil {
		return entity.Role{}, err
	}
	// This is a mock implementation
	if userID == "admin" {
		return entity.Role{ID: "1", Name: "Admin"}, nil
	} else if userID == "user" {
		return entity.Role{ID: "2", Name: "User"}, nil
	}
	return entity.Role{}, interfaces.ErrRoleNotFound
}

func (r *DB) UpdateRole(ctx context.Context, userID string, newRole entity.Role) error {
	if userID == "" || newRole.Name == "" {
		return errors.New("invalid input")
	}

	query := "UPDATE roles SET role = $1 WHERE user_id = $2"
	if _, err := r.conn.Conn().ExecContext(ctx, query, newRole.Name, userID); err != nil {
		return fmt.Errorf("failed to update role: %w", err)
	}
	return nil
}
