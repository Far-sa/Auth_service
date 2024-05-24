package main

import (
	"database/sql"
	"log"
	"net"
	"order-service/di"
	"order-service/internal/infrastructure/database"
	pb "order-service/pb"
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

	handler, err := di.InitializeOrderHandler(databaseURL, rabbitMQURL, rabbitMQExchange)
	if err != nil {
		log.Fatalf("failed to initialize order handler: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, handler)

	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("failed to listen on port 50055: %v", err)
	}

	log.Println("Order Service is running on port 50055")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
