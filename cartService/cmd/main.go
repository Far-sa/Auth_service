package main

import (
	"cart-service/di"
	"cart-service/internal/infrastructure/database"
	pb "cart-service/pb"
	"database/sql"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	rabbitMQExchange := os.Getenv("RABBITMQ_EXCHANGE")

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	err = database.Migrate(db)
	if err != nil {
		log.Fatalf("failed to migrate the database: %v", err)
	}

	handler, err := di.InitializeCartHandler(databaseURL, rabbitMQURL, rabbitMQExchange)
	if err != nil {
		log.Fatalf("failed to initialize cart handler: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCartServiceServer(grpcServer, handler)

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen on port 50054: %v", err)
	}

	log.Println("Cart Service is running on port 50054")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
