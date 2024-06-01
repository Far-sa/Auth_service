// cmd/http_server/main.go
package main

// import (
// 	httpServer "authentication-service/delivery/http"
// 	"authentication-service/domain/services"
// 	"authentication-service/infrastructure/database"
// 	"authentication-service/infrastructure/database/migrator"
// 	"authentication-service/infrastructure/messaging"
// 	"authentication-service/infrastructure/repository"
// 	"log"
// 	"net/http"
// 	"path"

// 	standard_runtime "runtime"
// )

// // Dummy implementation of UserRepository

// func main() {

// 	dsn := "postgres://postgres:password@localhost:5432/auth_db?sslmode=disable"
// 	db, err := database.NewSQLDB(dsn)
// 	if err != nil {
// 		log.Fatalf("Failed to create database: %v", err)
// 	}

// 	_, filename, _, _ := standard_runtime.Caller(0)
// 	dir := path.Join(path.Dir(filename), "../infrastructure/database/migrations")
// 	// Create a new migrator instance.
// 	migrator, err := migrator.NewMigrator(db.Conn(), dir)
// 	if err != nil {
// 		log.Fatalf("Failed to create migrator: %v", err)
// 	}

// 	// Apply all up migrations.
// 	if err := migrator.Up(); err != nil {
// 		log.Fatalf("Failed to migrate up: %v", err)
// 	}
// 	userRepo := repository.NewRepository(db)

// 	amqpUrl := "amqp://guest:guest@rabbitmq:5672/"
// 	publisher, _ := messaging.NewRabbitMQ(amqpUrl)
// 	authService := services.NewAuthService(userRepo, publisher)

// 	authHandler := httpServer.NewHTTPAuthHandler(authService)

// 	http.HandleFunc("/login", authHandler.Login)

// 	log.Println("HTTP server is running on port 8080")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
