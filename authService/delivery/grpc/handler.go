package grpcServer

import (
	"authentication-service/interfaces"
	auth "authentication-service/pb"
	mapper "authentication-service/utils/protobufMapper"
	"context"
	// Replace with your package paths
)

type grpcHandler struct {
	authService interfaces.AuthenticationService // Interface for dependency injection
	auth.UnimplementedAuthServiceServer
}

func NewGRPCHandler(authService interfaces.AuthenticationService) *grpcHandler {
	return &grpcHandler{authService: authService}
}

// Implement handler functions for other gRPC service methods defined in your `auth.proto` file
func (s *grpcHandler) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {

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

// func (s *grpcServer) Serve() {

// 	// listener
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50054))
// 	if err != nil {
// 		panic(err)
// 	}

// 	authServiceServer := grpcServer{}
// 	grpcServer := grpc.NewServer()
// 	pb.RegisterAuthServiceServer(grpcServer, authServiceServer) // Register handler for AuthService

// 	// Enable server reflection
// 	reflection.Register(grpcServer)
// 	// Serve
// 	go func() {
// 		if err := grpcServer.Serve(lis); err != nil {
// 			log.Fatalf("Failed to serve gRPC server: %v", err)
// 		}
// 	}()
// }
