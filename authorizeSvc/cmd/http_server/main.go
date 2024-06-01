package main

// import (
// 	httpHandler "authorization-service/delivery/http"
// 	"authorization-service/infrastructure/database"
// 	"authorization-service/infrastructure/database/migrator"
// 	"authorization-service/infrastructure/messaging"
// 	"authorization-service/infrastructure/messaging/rabbitmq"
// 	"authorization-service/infrastructure/repository"
// 	"authorization-service/internal/service"
// 	"path"
// 	standard_runtime "runtime"

// 	"context"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	"github.com/labstack/echo/v4"
// )

// func main() {

// 	//TODO add config to load them from environment variable

// 	dsn := "postgres://postgres:password@localhost:5432/authz_db?sslmode=disable"
// 	db, err := database.NewSQLDB(dsn)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
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

// 	// Initialize repository, service, and handler
// 	userRepo := repository.NewRepository(db)

// 	amqpUrl := "amqp://guest:guest@rabbitmq:5672/"
// 	rabbitAdapter, err := rabbitmq.NewRabbitMQAdapter(amqpUrl)
// 	if err != nil {
// 		log.Fatalf("Failed to create RabbitMQ adapter: %v", err)
// 	}
// 	defer rabbitAdapter.Close()

// 	consumer, err := messaging.NewRabbitMQConsumer(rabbitAdapter, "user_authenticated_queue", "user.authenticated", "auth_exchange", nil)
// 	if err != nil {
// 		log.Fatalf("Failed to create RabbitMQ consumer: %v", err)
// 	}

// 	authzService := service.NewAuthorizationService(userRepo, consumer)

// 	authzHandler := httpHandler.NewHTTPAuthzHandler(authzService)

// 	e := echo.New()
// 	e.POST("/assign-role", authzHandler.AssignRole)

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	go func() {
// 		log.Println("HTTP server is running on port 8081")
// 		if err := e.Start(":8081"); err != nil {
// 			log.Fatalf("failed to serve: %v", err)
// 		}
// 	}()

// 	// Start the authorization service

// 	// Handle OS signals for graceful shutdown
// 	signalChan := make(chan os.Signal, 1)
// 	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

// 	select {
// 	case <-ctx.Done():
// 		log.Println("Shutdown initiated")
// 	case sig := <-signalChan:
// 		log.Printf("Received signal: %v. Shutting down...", sig)
// 		cancel()
// 	}

// 	// Perform any cleanup tasks here if necessary
// 	log.Println("Server gracefully stopped")
// }
