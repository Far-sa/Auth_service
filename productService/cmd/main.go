package main

import (
	"database/sql"
	"log"
	"net"
	"os"
	"product-service/di"
	"product-service/internal/infrastructure/database"
	pb "product-service/pb"

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

	handler, err := di.InitializeProductHandler(databaseURL, rabbitMQURL, rabbitMQExchange)
	if err != nil {
		log.Fatalf("failed to initialize product handler: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, handler)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen on port 50053: %v", err)
	}

	log.Println("Product Catalog Service is running on port 50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
