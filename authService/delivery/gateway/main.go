package main

import (
	auth "authentication-service/pb"
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func RunHTTPGateway(ctx context.Context, grpcEndpoint string) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := auth.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts); err != nil {
		return err
	}

	log.Println("Serving gRPC-Gateway on http://localhost:8090")
	return http.ListenAndServe(":8080", mux)
}
