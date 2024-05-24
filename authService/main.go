package main

import (
	handler "authentication-service/delivery/grpc"
	"authentication-service/domain/services"
	"authentication-service/infrastructure/database"
	"authentication-service/infrastructure/messaging"
)

func main() {

	// create database connection

	// Create UserRepository and AuthService instances
	userRepository := database.NewPostgresUserRepository(db)
	rabbitUrl := ""
	messagePublisher, _ := messaging.NewRabbitMQPublisher(rabbitUrl)
	authService := services.NewAuthService(userRepository, messagePublisher)

	handler.NewAuthHandler(authService)

}
