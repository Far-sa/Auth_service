//go:build wireinject
// +build wireinject

package di

import (
	"authentication-service/interfaces"
	"authentication-service/internal/app/handler"
	"authentication-service/internal/app/service"
	"authentication-service/internal/infrastructure/database"
	"authentication-service/internal/infrastructure/messaging"
	"authentication-service/internal/infrastructure/repository"
	"database/sql"

	"github.com/google/wire"
	"github.com/streadway/amqp"
)

// ProvideDatabase returns a new instance of the database connection.
func ProvideDatabase() (*sql.DB, error) {
	db, err := database.NewPostgresDB("postgres://user:password@localhost/authorization?sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = database.Migrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ProvideRabbitMQConnection returns a new instance of the RabbitMQ connection.
func ProvideRabbitMQConnection() (*messaging.RabbitMQConnection, func(), error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return nil, nil, err
	}

	err = messaging.SetupRabbitMQ(conn)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		conn.Close()
	}
	return conn, cleanup, nil
}

// ProvideUserRepository returns a new instance of the user repository.
func ProvideUserRepository(db *sql.DB) interfaces.UserRepository {
	return repository.NewPostgresUserRepository(db)
}

// ProvideEventPublisher returns a new instance of the event publisher.
func ProvideEventPublisher(conn *messaging.RabbitMQConnection) interfaces.EventPublisher {
	return messaging.NewRabbitMQPublisher(conn)
}

// ProvideAuthService returns a new instance of the auth service.
func ProvideAuthService(repo interfaces.UserRepository, jwtKey []byte, publisher interfaces.EventPublisher) *service.AuthService {
	return service.NewAuthService(repo, jwtKey, publisher)
}

// ProvideAuthHandler returns a new instance of the auth handler.
func ProvideAuthHandler(service *service.AuthService) *handler.AuthHandler {
	return handler.NewAuthHandler(service)
}

// InitializeAuthHandler initializes the AuthHandler with all dependencies.
func InitializeAuthHandler(jwtKey []byte) (*handler.AuthHandler, func(), error) {
	wire.Build(
		ProvideDatabase,
		ProvideRabbitMQConnection,
		ProvideUserRepository,
		ProvideEventPublisher,
		ProvideAuthService,
		ProvideAuthHandler,
	)
	return &handler.AuthHandler{}, nil, nil
}
