package interfaces

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AuthorizationService interface {
	AssignRole(ctx context.Context, userID string) error
	// CheckPermission(ctx context.Context, username, permission string) (bool, error)
	// HandleUserAuthenticatedEvent(ctx context.Context, message string) error
}

type RoleRepository interface {
	AssignRole(ctx context.Context, userID, role string) error
	CheckPermission(ctx context.Context, username, permission string) (bool, error)
	UpdateUserRoles(ctx context.Context, userID string, role string) error
}

type RoleEvents interface {
	DeclareExchange(name, kind string) error
	CreateQueue(queueName string, durable, autodelete bool) (amqp.Queue, error)
	CreateBinding(queueName, routingKey, exchangeName string) error
	Consume(queueName, consumer string, autoAck bool) (<-chan amqp.Delivery, error)
	Publish(ctx context.Context, exchange, routingKey string, options amqp.Publishing) error
}
