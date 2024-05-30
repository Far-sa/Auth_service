package delivery

import (
	"authorization-service/internal/interfaces"
	"authorization-service/pb"
	"context"
	"fmt"
	"log"
	"net"

	// Replace with your package paths
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	authService interfaces.AuthorizationService // Interface for dependency injection
	pb.UnimplementedAuthorizationServiceServer
}

func NewGRPCServer(authService interfaces.AuthorizationService) *grpcServer {
	return &grpcServer{authService: authService}
}

func (s *grpcServer) Serve() {

	// listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		panic(err)
	}

	authorizationServiceServer := grpcServer{}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthorizationServiceServer(grpcServer, authorizationServiceServer) // Register handler for AuthService

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
func (s grpcServer) AssignRole(ctx context.Context, req *pb.AssignRoleRequest) (*pb.AssignRoleResponse, error) {
	err := s.authService.AssignRole(ctx, req.Username, req.Role)
	if err != nil {
		return nil, err
	}
	return &pb.AssignRoleResponse{Message: "Role assigned successfully"}, nil
}

func (s grpcServer) CheckPermission(ctx context.Context, req *pb.CheckPermissionRequest) (*pb.CheckPermissionResponse, error) {
	hasPermission, err := s.authService.CheckPermission(ctx, req.Username, req.Permission)
	if err != nil {
		return nil, err
	}
	//return &pb.CheckPermissionResponse{Has_permission: hasPermission}, nil
	return &pb.CheckPermissionResponse{HasPermission: hasPermission}, nil
}
