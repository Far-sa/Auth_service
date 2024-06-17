package main

import (
	"context"
	"log"
	"net/http"

	auth "authentication-service/auth"
	authz "authorization-service/authz"
	user "user-service/user"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func RunHTTPGateway(ctx context.Context, grpcEndpoint string) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, opts); err != nil {
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
