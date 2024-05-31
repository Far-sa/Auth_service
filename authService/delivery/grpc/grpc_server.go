package delivery

import (
	"authentication-service/interfaces"
	"authentication-service/pb"
	mapper "authentication-service/utils/protobufMapper"
	"context"
	"fmt"
	"log"
	"net"

	// Replace with your package paths
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	authService interfaces.AuthenticationService // Interface for dependency injection
	pb.UnimplementedAuthServiceServer
}

func NewGRPC(authService interfaces.AuthenticationService) *grpcServer {
	return &grpcServer{authService: authService}
}

func (s *grpcServer) Serve() {

	// listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50054))
	if err != nil {
		panic(err)
	}

	authServiceServer := grpcServer{}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authServiceServer) // Register handler for AuthService

	// Enable server reflection
	reflection.Register(grpcServer)
	// Serve
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
}

// Implement handler functions for other gRPC service methods defined in your `auth.proto` file
func (s grpcServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	//* convert request from protobuf to param.LoginRequest
	paramReq := mapper.PbToParamLoginRequest(req)

	resp, err := s.authService.Login(ctx, paramReq)
	if err != nil {
		return nil, err
	}

	//* converts the result DTOs back to Protobuf messages
	protoResp := mapper.ParamToPbLoginResponse(resp)

	return protoResp, nil
}
