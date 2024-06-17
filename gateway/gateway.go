package main

import (
	"context"
	"log"
	"net/http"

	auth "gateway/proto/auth"
	authz "gateway/proto/authz"
	user "gateway/proto/user"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

//TODO modify protoc generate path to avoid code duplication

func RunHTTPGateway(ctx context.Context, grpcEndpoint string) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts); err != nil {
		return err
	}

	if err := auth.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts); err != nil {
		return err
	}

	if err := authz.RegisterAuthorizationServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts); err != nil {
		return err
	}

	log.Println("Serving gRPC-Gateway on http://localhost:8081")
	return http.ListenAndServe(":8081", mux)
}

func main() {
	// Start gRPC-Gateway in a separate goroutine
	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		if err := RunHTTPGateway(ctx, ":50051"); err != nil {
			log.Fatalf("Failed to run gRPC-Gateway: %v", err)
		}
	}()

}
