//go:build wireinject
// +build wireinject

package di

import (
	"database/sql"
	"order-service/interfaces"
	"order-service/internal/app/handler"
	"order-service/internal/app/service"
	"order-service/internal/infrastructure/database"
	"order-service/internal/infrastructure/messaging"

	"github.com/google/wire"
)

func InitializeOrderHandler(databaseURL, rabbitMQURL, rabbitMQExchange string) (*handler.OrderHandler, error) {
	wire.Build(
		NewDatabase,
		NewRabbitMQ,
		database.NewPostgresOrderRepository,
		service.NewOrderService,
		handler.NewOrderHandler,
		wire.Bind(new(interfaces.OrderRepository), new(*database.PostgresOrderRepository)),
		wire.Bind(new(interfaces.Messaging), new(*messaging.RabbitMQ)),
	)
	return &handler.OrderHandler{}, nil
}

func NewDatabase(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func NewRabbitMQ(rabbitMQURL, rabbitMQExchange string) (*messaging.RabbitMQ, error) {
	return messaging.NewRabbitMQ(rabbitMQURL, rabbitMQExchange)
}
