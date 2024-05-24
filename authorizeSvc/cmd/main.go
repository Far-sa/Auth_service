package main

import (
	"authorization-service/internal/di"
	"authorization-service/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	defaultPort = "50052"
)

func main() {

	// conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	// if err != nil {
	// 	log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	// }
	// defer conn.Close()

	// consumer, err := messaging.NewRabbitMQConsumer(conn)
	// if err != nil {
	// 	log.Fatalf("Failed to create RabbitMQ consumer: %v", err)
	// }

	// Start consuming messages
	//defer consumer.channel.Close()
	// Initialize Authz Handler with DI
	authzHandler, cleanup, err := di.InitializeAuthzHandler()
	if err != nil {
		log.Fatalf("Failed to initialize authz handler: %v", err)
	}
	defer cleanup()

	listener, err := net.Listen("tcp", ":"+defaultPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", defaultPort, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthorizationServiceServer(grpcServer, authzHandler)

	log.Println("Starting gRPC server on port", defaultPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
