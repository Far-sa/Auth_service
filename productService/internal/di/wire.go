//go:build wireinject
// +build wireinject

package di

import (
	"database/sql"
	"product-service/interfaces"
	"product-service/internal/app/handler"
	"product-service/internal/app/service"
	"product-service/internal/infrastructure/messaging"
	"product-service/internal/infrastructure/repository"

	"github.com/google/wire"
)

func InitializeProductHandler(db *sql.DB, rabbitMQURL, rabbitMQExchange string) (*handler.ProductHandler, error) {
	wire.Build(
		repository.NewPostgresProductRepository,
		service.NewProductService,
		handler.NewProductHandler,
		messaging.NewRabbitMQ,
		wire.Bind(new(interfaces.ProductRepository), new(*repository.PostgresProductRepository)),
	)
	return &handler.ProductHandler{}, nil
}
