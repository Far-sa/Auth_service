package interfaces

import (
	"context"
	"user-service/internal/entity"
	"user-service/internal/param"

	amqp "github.com/rabbitmq/amqp091-go"
)

type UserService interface {
	GetUser(userID string) (param.UserProfileResponse, error)
	GetUserByEmail(ctx context.Context, email string) (param.UserProfileResponse, error)
}

type UserRepository interface {
	GetUserByID(userID string) (entity.UserProfile, error)
	FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*entity.UserProfile, error)
}
type UserEvents interface {
	DeclareExchange(name, kind string) error
	CreateQueue(queueName string, durable, autodelete bool) (amqp.Queue, error)
	CreateBinding(queueName, routingKey, exchangeName string) error
	Consume(queueName, consumer string, autoAck bool) (<-chan amqp.Delivery, error)
	Publish(ctx context.Context, exchange, routingKey string, options amqp.Publishing) error
}
