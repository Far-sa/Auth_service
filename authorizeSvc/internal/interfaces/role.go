package interfaces

import (
	"authorization-service/internal/entity"
	"authorization-service/internal/param"
	"context"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

var ErrRoleNotFound = errors.New("role not found")

type AuthorizationService interface {
	AssignRole(ctx context.Context, userID param.RoleAssignmentRequest) error
	UpdateUserRole(ctx context.Context, userID string, newRole param.RoleUpdateRequest) error

	// CheckPermission(ctx context.Context, username, permission string) (bool, error)
	// HandleUserAuthenticatedEvent(ctx context.Context, message string) error
}

type RoleRepository interface {
	CheckPermission(ctx context.Context, username, permission string) (bool, error)
	// UpdateUserRoles(ctx context.Context, userID string, role string) error
	//!
	AssignRole(ctx context.Context, userID, role string) error
	GetRoleByUserID(ctx context.Context, userID string) (entity.Role, error)
	UpdateRole(ctx context.Context, userID string, newRole entity.Role) error
}

type RoleEvents interface {
	DeclareExchange(name, kind string) error
	CreateQueue(queueName string, durable, autodelete bool) (amqp.Queue, error)
	CreateBinding(queueName, routingKey, exchangeName string) error
	Consume(queueName, consumer string, autoAck bool) (<-chan amqp.Delivery, error)
	Publish(ctx context.Context, exchange, routingKey string, options amqp.Publishing) error
}
