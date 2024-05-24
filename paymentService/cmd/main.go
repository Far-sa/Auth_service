package main

import (
	"database/sql"
	"log"
	"net"
	"os"
	"payment-service/di"
	"payment-service/internal/infrastructure/database"
	pb "payment-service/pb"

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
	paymentHandler, err := di.InitializePaymentHandler(db, rabbitMQUrl)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, paymentHandler)

	log.Println("Starting Payment Service on port :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
