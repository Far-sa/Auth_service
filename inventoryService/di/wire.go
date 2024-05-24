//go:build wireinject
// +build wireinject

package di

import (
	"database/sql"
	"inventory-service/interfaces"
	"inventory-service/internal/app/handler"
	"inventory-service/internal/app/service"
	"inventory-service/internal/infrastructure/database"
	"inventory-service/internal/infrastructure/messaging"

	"github.com/google/wire"
)

func InitializeInventoryHandler(db *sql.DB, rabbitMQUrl string) (*handler.InventoryHandler, error) {
	wire.Build(
		database.NewInventoryRepository,
		messaging.NewRabbitMQ,
		service.NewInventoryService,
		handler.NewInventoryHandler,
		wire.Bind(new(interfaces.InventoryRepository), new(*database.InventoryRepository)),
		wire.Bind(new(interfaces.Messaging), new(*messaging.RabbitMQ)),
	)
	return &handler.InventoryHandler{}, nil
}
