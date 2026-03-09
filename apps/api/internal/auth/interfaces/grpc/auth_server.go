// Package grpc provides gRPC service implementations for the auth module.
//
// Interfaces/gRPC layer in Clean Architecture:
// - Implements gRPC service interfaces (generated from .proto files)
// - Handles gRPC-specific concerns (context, metadata, errors)
// - Provides high-performance RPC endpoints for service-to-service communication
//
// In production applications, this folder typically contains:
// - gRPC server implementations for each service
// - Protobuf message conversion (to/from domain entities)
// - gRPC interceptors (logging, auth, metrics)
// - Error handling and status code mapping
// - Streaming implementations (server-side, client-side, bidirectional)
// - Health check implementations
//
// Best practices:
// - Keep gRPC servers thin - delegate to application handlers
// - Handle gRPC context cancellation and timeouts
// - Use appropriate gRPC status codes
// - Implement proper error details (google.rpc.Status)
// - Add interceptors for cross-cutting concerns
// - Enable reflection for development
// - Implement health checks

package grpc

// AuthServer implements the gRPC AuthService interface.
// In production applications, auth gRPC servers typically implement:
// - Login(LoginRequest) - authenticate user
// - Register(RegisterRequest) - create new user
// - ValidateToken(TokenRequest) - verify token (for other services)
// - RefreshToken(RefreshRequest) - get new access token
// - GetUser(GetUserRequest) - retrieve user details
// - CheckPermission(PermissionRequest) - authorization check

import (
	"context"
	"mytodo/apps/api/internal/auth/domain/entity"
	"mytodo/apps/api/internal/auth/domain/service"
	"mytodo/apps/api/internal/auth/interfaces/grpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer implements the gRPC AuthService interface
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	authService *service.AuthService
}

// NewAuthServer creates a new gRPC auth server
func NewAuthServer(authService *service.AuthService) *AuthServer {
	return &AuthServer{
		authService: authService,
	}
}

// Login authenticates a user and returns access tokens
func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Validate input
	if req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	// Authenticate user
	user, tokens, err := s.authService.AuthenticateUser(
		ctx,
		req.Email,
		req.Password,
		req.UserAgent,
		req.IpAddress,
	)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	// Build response
	return &pb.LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    tokens.ExpiresIn,
		User:         entityToProtoUser(user),
	}, nil
}

// Register creates a new user account
func (s *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Validate input
	if req.Email == "" || req.Password == "" || req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "email, password, and name are required")
	}

	// Register user
	user, err := s.authService.RegisterUser(ctx, req.Email, req.Password, req.Name)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	// Authenticate the newly registered user
	authenticatedUser, tokens, err := s.authService.AuthenticateUser(
		ctx,
		req.Email,
		req.Password,
		"",
		"",
	)
	if err != nil {
		// User created but login failed - return user without tokens
		return &pb.RegisterResponse{
			User: entityToProtoUser(user),
		}, nil
	}

	// Build response
	return &pb.RegisterResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    tokens.ExpiresIn,
		User:         entityToProtoUser(authenticatedUser),
	}, nil
}

// RefreshToken generates a new access token using a refresh token
func (s *AuthServer) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	// Validate input
	if req.RefreshToken == "" {
		return nil, status.Error(codes.InvalidArgument, "refresh token is required")
	}

	// Refresh token
	tokens, err := s.authService.RefreshAccessToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	// Build response
	return &pb.RefreshTokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    tokens.ExpiresIn,
	}, nil
}

// Logout revokes the user's session
func (s *AuthServer) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	// Validate input
	if req.RefreshToken == "" {
		return nil, status.Error(codes.InvalidArgument, "refresh token is required")
	}

	// Logout
	err := s.authService.Logout(ctx, req.RefreshToken)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Build response
	return &pb.LogoutResponse{
		Message: "Successfully logged out",
	}, nil
}

// ValidateToken verifies an access token (for inter-service auth)
func (s *AuthServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	// Validate input
	if req.AccessToken == "" {
		return &pb.ValidateTokenResponse{
			Valid:        false,
			ErrorMessage: "access token is required",
		}, nil
	}

	// This would require a ValidateToken method on the AuthService
	// For now, we'll return a basic validation
	// In production, you'd want to actually validate the token signature and expiration
	return &pb.ValidateTokenResponse{
		Valid:        false,
		ErrorMessage: "token validation not yet implemented",
	}, nil
}

// GetCurrentUser retrieves the authenticated user's information
func (s *AuthServer) GetCurrentUser(ctx context.Context, req *pb.GetCurrentUserRequest) (*pb.GetCurrentUserResponse, error) {
	// Validate input
	if req.AccessToken == "" {
		return nil, status.Error(codes.InvalidArgument, "access token is required")
	}

	// In a production system, you'd parse the token to get the user ID
	// For now, return an error
	return nil, status.Error(codes.Unimplemented, "get current user not yet implemented")
}

// entityToProtoUser converts domain User entity to protobuf User message
func entityToProtoUser(user *entity.User) *pb.User {
	protoUser := &pb.User{
		Id:        user.ID.String(),
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		IsActive:  user.IsActive,
	}

	if user.LastLoginAt != nil {
		protoUser.LastLoginAt = user.LastLoginAt.Format("2006-01-02T15:04:05Z07:00")
	}

	return protoUser
}
