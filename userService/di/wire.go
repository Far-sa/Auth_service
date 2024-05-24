// +build wireinject

package di

import (
    "database/sql"
    "user-service/internal/app/handler"
    "user-service/internal/app/service"
    "user-service/internal/infrastructure/database"
    "user-service/internal/infrastructure/messaging"
    "user-service/interfaces"
    "github.com/google/wire"
)

func InitializeUserHandler(db *sql.DB, rabbitMQUrl string) (*handler.UserHandler, error) {
    wire.Build(
        database.NewUserRepository,
        messaging.NewRabbitMQ,
        service.NewUserService,
        handler.NewUserHandler,
        wire.Bind(new(interfaces.UserRepository), new(*database.UserRepository)),
        wire.Bind(new(interfaces.Messaging), new(*messaging.RabbitMQ)),
    )
    return &handler.UserHandler{}, nil
}
