package delivery_test

import (
	delivery "authorization-service/delivery/gprc"
	"authorization-service/pb"
	"context"
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/test/bufconn"
)

// MockAuthorizationService is a mock implementation of interfaces.AuthorizationService
type MockAuthorizationService struct {
	mock.Mock
}

// MockAuthzHandler is a mock implementation of handler.AuthzHandler
type MockAuthzHandler struct {
	mock.Mock
	pb.UnimplementedAuthorizationServiceServer // Embedding this is important to satisfy the interface requirements
}

func (m *MockAuthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func TestNewGRPCServer(t *testing.T) {
	mockAuthService := new(MockAuthorizationService)
	mockAuthHandler := new(MockAuthzHandler)

	server := delivery.NewGRPCServer(mockAuthService, mockAuthHandler)
	assert.NotNil(t, server)
	assert.Equal(t, mockAuthService, server.authService)
	assert.Equal(t, mockAuthHandler, server.authHandler)
}

func TestServe(t *testing.T) {
	mockAuthService := new(MockAuthorizationService)
	mockAuthHandler := new(MockAuthzHandler)

	server := delivery.NewGRPCServer(mockAuthService, mockAuthHandler)

	// Create an in-memory listener for testing
	lis := bufconn.Listen(1024 * 1024)

	go func() {
		if err := server.Serve(lis); err != nil {
			t.Errorf("Server exited with error: %v", err)
		}
	}()

	// Dial the in-memory listener
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	// Check if reflection service is available
	client := reflection.NewServerReflectionClient(conn)
	reflectionTestCall(client, t)

	// Additional checks can be made here to ensure that the authHandler is correctly registered and serving requests
}

func reflectionTestCall(client reflection.ServerReflectionClient, t *testing.T) {
	// Perform server reflection call to test if the server is running and reflection is enabled
	stream, err := client.ServerReflectionInfo(context.Background())
	if err != nil {
		t.Fatalf("Failed to call ServerReflectionInfo: %v", err)
	}
	_, err = stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive from ServerReflectionInfo stream: %v", err)
	}
}
