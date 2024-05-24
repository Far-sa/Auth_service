//go:build wireinject
// +build wireinject

package di

import (
	"cart-service/interfaces"
	"cart-service/internal/app/service"
	"cart-service/internal/infrastructure/database"
	"cart-service/internal/infrastructure/messaging"
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

//! provide repository,service and handler

func InitializeCartHandler(databaseURL, rabbitMQURL, rabbitMQExchange string) (*handler.CartHandler, error) {
	wire.Build(
		NewDatabase,
		NewRabbitMQ,
		database.NewPostgresCartRepository,
		service.NewCartService,
		handler.NewCartHandler,
		wire.Bind(new(interfaces.CartRepository), new(*database.PostgresCartRepository)),
		wire.Bind(new(interfaces.Messaging), new(*messaging.RabbitMQ)),
	)
	return &handler.CartHandler{}, nil
}
