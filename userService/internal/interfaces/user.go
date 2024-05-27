package interfaces

import (
	"context"
	"user-service/internal/entity"
	"user-service/internal/param"

	"github.com/streadway/amqp"
)

type UserService interface {
	// CreateUser(user entity.UserDetail) error
	GetUser(userID string) (param.UserInfo, error)
}

type UserRepository interface {
	// CreateUser(user entity.UserDetail) error
	GetUser(userID string) (entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error

	FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*entity.User, error)
}
type Messaging interface {
	Publish(message []byte, queueName string) error
	Consume(queue string) (<-chan amqp.Delivery, error)

	//Subscribe(queueName string, handler func(message []byte)) error
}
