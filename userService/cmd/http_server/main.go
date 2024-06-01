package httpServer

import (
	"authentication-service/infrastructure/database/migrator"
	"log"
	"path"
	standard_runtime "runtime"

	httpHandler "user-service/delivery/http"
	"user-service/infrastructure/database"
	"user-service/infrastructure/messaging"
	"user-service/infrastructure/repository"
	"user-service/internal/service"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func HttpServer() {

	dsn := "postgres://postgres:password@localhost:5432/user_db?sslmode=disable"
	db, err := database.NewSQLDB(dsn)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	//! runtime.Caller(0)` returns the file name and line number of the caller's caller.
	//! `path.Dir(filename)` returns the directory of the `main.go` file. `path.Join` constructs the path to the migrations directory.
	_, filename, _, _ := standard_runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../infrastructure/database/migrations")
	// Create a new migrator instance.
	migrator, err := migrator.NewMigrator(db.Conn(), dir)
	if err != nil {
		log.Fatalf("Failed to create migrator: %v", err)
	}

	// Apply all up migrations.
	if err := migrator.Up(); err != nil {
		log.Fatalf("Failed to migrate up: %v", err)
	}
	// Initialize repository, service, and handler
	userRepo := repository.NewRepository(db)
	amqpUrl := "amqp://guest:guest@localhost:5672/"
	userEvent, _ := messaging.NewRabbitMQ(amqpUrl)
	userSvc := service.NewUserService(userRepo, userEvent)

	userHandler := httpHandler.NewHTTPAuthHandler(userSvc)

	e := echo.New()
	e.POST("/register", userHandler.SignUp)
	e.GET("/getUser", userHandler.GetUserByEmail)

	log.Println("HTTP server is running on port 8081")
	if err := e.Start(":8081"); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
