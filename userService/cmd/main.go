package main

import (
	"log"
	"net"
	"path"
	standard_runtime "runtime"

	grpcHandler "user-service/delivery/grpc"
	httpHandler "user-service/delivery/http"
	"user-service/infrastructure/database"
	"user-service/infrastructure/database/migrator"
	"user-service/infrastructure/messaging"
	"user-service/infrastructure/repository"
	"user-service/internal/service"
	user "user-service/pb"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {

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

	// Start gRPC server in a separate goroutine
	go func() {
		userHandler := grpcHandler.New(userSvc)

		grpcServer := grpc.NewServer()
		user.RegisterUserServiceServer(grpcServer, userHandler)

		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// Start gRPC-Gateway in a separate goroutine
	// go func() {
	// 	ctx := context.Background()
	// 	ctx, cancel := context.WithCancel(ctx)
	// 	defer cancel()

	// 	if err := gateway.RunHTTPGateway(ctx, ":50051"); err != nil {
	// 		log.Fatalf("Failed to run gRPC-Gateway: %v", err)
	// 	}
	// }()

	// Start HTTP server in the main goroutine
	userHandler := httpHandler.NewHTTPAuthHandler(userSvc)

	e := echo.New()
	e.POST("/register", userHandler.SignUp)
	e.GET("/getUser/:id", userHandler.GetUser)

	log.Println("HTTP server is running on port 8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
