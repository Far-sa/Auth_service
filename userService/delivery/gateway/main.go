package gateway

import (
	"context"
	"log"
	"net/http"
	"user-service/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// func main() {
// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()

// 	mux := runtime.NewServeMux()
// 	opts := []grpc.DialOption{grpc.WithInsecure()}
// 	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:50053", opts)
// 	if err != nil {
// 		log.Fatalf("Failed to register gRPC-Gateway: %v", err)
// 	}

// 	log.Println("Serving gRPC-Gateway on http://localhost:8080")
// 	if err := http.ListenAndServe(":8080", mux); err != nil {
// 		log.Fatalf("Failed to serve gRPC-Gateway: %v", err)
// 	}
// }

func RunHTTPGateway(ctx context.Context, grpcEndpoint string) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts); err != nil {
		return err
	}

	log.Println("Serving gRPC-Gateway on http://localhost:8080")
	return http.ListenAndServe(":8080", mux)
}
