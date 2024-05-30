package interfaces

import (
	"authentication-service/domain/entities"
	"authentication-service/domain/param"
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AuthenticationService interface {
	Login(ctx context.Context, loginRequest param.LoginRequest) (param.LoginResponse, error)
	Register(ctx context.Context, req param.RegisterRequest) error
	// ... other authentication methods (e.g., refresh token)
}

type AuthRepository interface {
	SaveToke(ctx context.Context, user *entities.TokenPair) error
	FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*entities.User, error)
	SaveUser(ctx context.Context, user entities.User) error
}

type AuthEvents interface {
	DeclareExchange(name, kind string) error
	CreateQueue(queueName string, durable, autodelete bool) (amqp.Queue, error)
	CreateBinding(queueName, routingKey, exchangeName string) error
	Consume(queueName, consumer string, autoAck bool) (<-chan amqp.Delivery, error)
	Publish(ctx context.Context, exchange, routingKey string, options amqp.Publishing) error
}
