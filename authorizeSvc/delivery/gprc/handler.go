package grpcHandler

import (
	"authorization-service/internal/interfaces"
	authz "authorization-service/pb"
	"authorization-service/utils/mapper"
	"context"
	"fmt"
	// Replace with your package paths
)

type grpcHandler struct {
	authService interfaces.AuthorizationService // Interface for dependency injection
	authz.UnimplementedAuthorizationServiceServer
}

func NewGRPCHandler(authService interfaces.AuthorizationService) *grpcHandler {
	return &grpcHandler{authService: authService}
}

//

// Implement handler functions for other gRPC service methods defined in your `auth.proto` file
func (s grpcHandler) AssignRole(ctx context.Context, req *authz.AssignRoleRequest) (*authz.AssignRoleResponse, error) {

	//* proto to param
	paramReq := mapper.PbToParamAssignRoleRequest(req)
	err := s.authService.AssignRole(ctx, paramReq)
	if err != nil {
		return nil, fmt.Errorf("failed to assign role: %w", err)
	}

	//* param to proto
	pbrRsp := mapper.ToPbAssignRoleResponse(&paramReq)
	return &authz.AssignRoleResponse{Message: pbrRsp.Message}, nil
}

func (s grpcHandler) UpdateRole(ctx context.Context, req *authz.UpdateRoleRequest) (*authz.UpdateRoleResponse, error) {

	paramReq := mapper.PbToParamUpdateRoleRequest(req)
	err := s.authService.UpdateUserRole(ctx, paramReq)
	if err != nil {
		return nil, fmt.Errorf("failed to update role: %w", err)
	}

	pbrRsp := mapper.ToPbUpdateRoleResponse(&paramReq)
	return &authz.UpdateRoleResponse{Message: pbrRsp.Message}, nil
}

// func (s grpcHandler) CheckPermission(ctx context.Context, req *authz.CheckPermissionRequest) (*authz.CheckPermissionResponse, error) {

// 	//*  proto to param
// 	hasPermission, err := s.authService.CheckPermission(ctx, req.Username, req.Permission)
// 	if err != nil {
// 		return nil, err
// 	}
// 	//return &pb.CheckPermissionResponse{Has_permission: hasPermission}, nil
// 	//* param to
// 	return &authz.CheckPermissionResponse{HasPermission: hasPermission}, nil
// }

// func (s *grpcServer) Serve() {

// 	// listener
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
// 	if err != nil {
// 		panic(err)
// 	}

// 	authorizationServiceServer := grpcServer{}

// 	grpcServer := grpc.NewServer()
// 	authz.RegisterAuthorizationServiceServer(grpcServer, authorizationServiceServer) // Register handler for AuthService

// 	// Enable server reflection
// 	reflection.Register(grpcServer)

// 	// Serve
// 	go func() {
// 		if err := grpcServer.Serve(lis); err != nil {
// 			log.Fatalf("Failed to serve gRPC server: %v", err)
// 		}
// 	}()
// }
