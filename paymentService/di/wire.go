//go:build wireinject
// +build wireinject

package di

import (
	"database/sql"
	"payment-service/interfaces"
	"payment-service/internal/app/handler"
	"payment-service/internal/app/service"
	"payment-service/internal/infrastructure/messaging"
	"payment-service/internal/infrastructure/repository"

	"github.com/google/wire"
)

func InitializePaymentHandler(db *sql.DB, rabbitMQUrl string) (*handler.PaymentHandler, error) {
	wire.Build(
		repository.NewPaymentRepository,
		messaging.NewRabbitMQ,
		service.NewPaymentService,
		handler.NewPaymentHandler,
		wire.Bind(new(interfaces.PaymentRepository), new(*database.PaymentRepository)),
		wire.Bind(new(interfaces.Messaging), new(*messaging.RabbitMQ)),
	)
	return &handler.PaymentHandler{}, nil
}
