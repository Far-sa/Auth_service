package main

import (
	"database/sql"
	"log"
	"net"
	"os"
	"user-service/di"
	"user-service/internal/infrastructure/database"
	pb "user-service/pb"

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
	userHandler, err := di.InitializeUserHandler(db, rabbitMQUrl)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	log.Println("Starting User Service on port :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
