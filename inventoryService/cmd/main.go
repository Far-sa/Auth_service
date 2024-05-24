package main

import (
	"database/sql"
	"inventory-service/di"
	"inventory-service/internal/infrastructure/database"
	pb "inventory-service/pb"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = database.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	rabbitMQUrl := os.Getenv("RABBITMQ_URL")
	inventoryHandler, err := di.InitializeInventoryHandler(db, rabbitMQUrl)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterInventoryServiceServer(grpcServer, inventoryHandler)

	log.Println("Starting Inventory Service on port :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
