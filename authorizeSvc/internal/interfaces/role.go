package interfaces

import "context"

type AuthorizationService interface {
	AssignRole(ctx context.Context, username, role string) error
	CheckPermission(ctx context.Context, username, permission string) (bool, error)
}

type RoleRepository interface {
	AssignRole(ctx context.Context, username, role string) error
	CheckPermission(ctx context.Context, username, permission string) (bool, error)
	UpdateUserRoles(ctx context.Context, userID string, role string) error
}

type EventPublisher interface {
	Publish(event string, data []byte) error
}
