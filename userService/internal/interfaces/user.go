package interfaces

import (
	"context"
	"user-service/internal/entity"
	"user-service/internal/param"

	amqp "github.com/rabbitmq/amqp091-go"
)

type UserService interface {
	GetUser(ctx context.Context, userID string) (param.UserProfileResponse, error)
	GetUserByEmail(ctx context.Context, email string) (param.UserProfileResponse, error)
}

type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*entity.UserProfile, error)
	FindUserByEmail(ctx context.Context, Email string) (*entity.UserProfile, error)
	CreateUser(ctx context.Context, user *entity.UserProfile) (*entity.UserProfile, error)
	// UpdateUser(ctx context.Context, user *entity.UserProfile) error
	// DeleteUser(ctx context.Context, userID string) error
}
type UserEvents interface {
	DeclareExchange(name, kind string) error
	CreateQueue(queueName string, durable, autodelete bool) (amqp.Queue, error)
	CreateBinding(queueName, routingKey, exchangeName string) error
	Consume(queueName, consumer string, autoAck bool) (<-chan amqp.Delivery, error)
	Publish(ctx context.Context, exchange, routingKey string, options amqp.Publishing) error
}
