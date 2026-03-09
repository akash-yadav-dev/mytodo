# gRPC Authentication Service

This document describes how to test the gRPC authentication service.

## Overview

The authentication service is now exposed via both HTTP (port 8080) and gRPC (port 50051).

### gRPC Service Definition

The service is defined in `apps/api/internal/auth/interfaces/grpc/proto/auth.proto` and includes:

- **Login** - Authenticate a user and return access tokens
- **Register** - Create a new user account
- **RefreshToken** - Generate new tokens using a refresh token
- **Logout** - Revoke a user's session
- **ValidateToken** - Verify an access token (for inter-service auth)
- **GetCurrentUser** - Retrieve authenticated user information

## Testing the gRPC Server

### Prerequisites

1. **Database**: Ensure PostgreSQL is running and configured
2. **Redis**: Ensure Redis is running and configured
3. **Environment Variables**: Set up required environment variables (JWT secret, DB connection, etc.)

### Method 1: Using the Test Client

We've created a simple test client that exercises all gRPC endpoints.

1. Start the server:
   ```bash
   cd apps/api/cmd/server
   ./server.exe
   ```

2. In another terminal, run the test client:
   ```bash
   cd apps/api/cmd/test-grpc-client
   go run main.go
   ```

The test client will:
- Register a new user
- Login with that user
- Refresh the access token
- Logout
- Attempt to validate a token

### Method 2: Using grpcurl

[grpcurl](https://github.com/fullstorydev/grpcurl) is a command-line tool for interacting with gRPC servers.

1. Install grpcurl:
   ```bash
   go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
   ```

2. List available services:
   ```bash
   grpcurl -plaintext localhost:50051 list
   ```

3. Describe the AuthService:
   ```bash
   grpcurl -plaintext localhost:50051 describe auth.AuthService
   ```

4. Register a new user:
   ```bash
   grpcurl -plaintext -d '{
     "email": "user@example.com",
     "password": "password123",
     "name": "Test User"
   }' localhost:50051 auth.AuthService/Register
   ```

5. Login:
   ```bash
   grpcurl -plaintext -d '{
     "email": "user@example.com",
     "password": "password123",
     "user_agent": "grpcurl",
     "ip_address": "127.0.0.1"
   }' localhost:50051 auth.AuthService/Login
   ```

6. Refresh token (use the refresh_token from login response):
   ```bash
   grpcurl -plaintext -d '{
     "refresh_token": "YOUR_REFRESH_TOKEN_HERE"
   }' localhost:50051 auth.AuthService/RefreshToken
   ```

7. Logout:
   ```bash
   grpcurl -plaintext -d '{
     "refresh_token": "YOUR_REFRESH_TOKEN_HERE"
   }' localhost:50051 auth.AuthService/Logout
   ```

### Method 3: Using BloomRPC or Postman

Both [BloomRPC](https://github.com/bloomrpc/bloomrpc) and [Postman](https://www.postman.com/) support gRPC:

1. Import the proto file: `apps/api/internal/auth/interfaces/grpc/proto/auth.proto`
2. Connect to `localhost:50051`
3. Call the methods interactively

## Architecture

The gRPC implementation follows Clean Architecture:

```
interfaces/grpc/
├── proto/
│   └── auth.proto          # Protocol Buffer definitions
├── pb/
│   ├── auth.pb.go          # Generated protobuf code
│   └── auth_grpc.pb.go     # Generated gRPC service code
└── auth_server.go          # gRPC server implementation
```

The gRPC server is thin and delegates to the domain service layer:
- `AuthServer` implements `pb.AuthServiceServer`
- Converts gRPC messages to/from domain entities
- Handles gRPC-specific error codes
- Uses the same `AuthService` as the HTTP layer

## Ports

- **HTTP REST API**: `http://localhost:8080`
- **gRPC API**: `localhost:50051`

Both use the same underlying authentication service for consistency.

## Next Steps

To fully implement the gRPC service, consider:

1. Implement `ValidateToken` by adding a validation method to `AuthService`
2. Implement `GetCurrentUser` by parsing JWT tokens
3. Add gRPC interceptors for:
   - Logging
   - Metrics
   - Authentication/Authorization
   - Rate limiting
4. Add TLS/SSL support for production
5. Implement streaming endpoints if needed
6. Add comprehensive integration tests
7. Set up API gateway if exposing to external clients

## Troubleshooting

### Server won't start
- Check if ports 8080 and 50051 are available
- Verify database and Redis connections
- Check environment variables

### Authentication fails
- Verify the database schema is properly initialized
- Check JWT secret configuration
- Ensure user exists in the database

### Connection refused
- Make sure the server is running
- Check firewall settings
- Verify you're connecting to the correct port (50051 for gRPC)
