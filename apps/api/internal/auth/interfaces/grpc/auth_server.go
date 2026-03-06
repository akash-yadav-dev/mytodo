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
