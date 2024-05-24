//go:build wireinject
// +build wireinject

package di

import (
	"authorization-service/interfaces"
	"authorization-service/internal/app/handler"
	"authorization-service/internal/app/service"
	"authorization-service/internal/infrastructure/database"
	"authorization-service/internal/infrastructure/messaging"
	"authorization-service/internal/infrastructure/repository"
	"database/sql"

	"github.com/google/wire"
	"github.com/streadway/amqp"
)

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

func ProvideRoleRepository(db *sql.DB) interfaces.RoleRepository {
	return repository.NewPostgresRoleRepository(db)
}

func ProvideEventPublisher(conn *messaging.RabbitMQConnection) interfaces.EventPublisher {
	return messaging.NewRabbitMQPublisher(conn)
}

func ProvideAuthzService(repo interfaces.RoleRepository, publisher interfaces.EventPublisher) *service.AuthzService {
	return service.NewAuthzService(repo, publisher)
}

func ProvideAuthzHandler(service *service.AuthzService) *handler.AuthzHandler {
	return handler.NewAuthzHandler(service)
}

func InitializeAuthzHandler() (*handler.AuthzHandler, func(), error) {
	wire.Build(
		ProvideDatabase,
		ProvideRabbitMQConnection,
		ProvideRoleRepository,
		ProvideEventPublisher,
		ProvideAuthzService,
		ProvideAuthzHandler,
	)
	return &handler.AuthzHandler{}, nil, nil
}
