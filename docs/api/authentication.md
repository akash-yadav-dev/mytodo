# Authentication API Documentation

## Overview

This document describes the authentication service implementation with JWT tokens and refresh token functionality for the MyTodo API.

## Features

- User registration with email and password
- User login with JWT access tokens
- JWT refresh token mechanism for obtaining new access tokens
- Secure logout with session revocation
- Password hashing with bcrypt
- Session management with database persistence
- Protected routes with JWT middleware

## Authentication Endpoints

### Base URL
```
http://localhost:8080/api/v1/auth
```

### 1. Register New User

**Endpoint:** `POST /api/v1/auth/register`

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "SecurePassword123"
}
```

**Response:** `201 Created`
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 3600,
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "john@example.com",
    "name": "John Doe"
  }
}
```

### 2. Login

**Endpoint:** `POST /api/v1/auth/login`

**Request Body:**
```json
{
  "email": "john@example.com",
  "password": "SecurePassword123"
}
```

**Response:** `200 OK`
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 3600,
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "john@example.com",
    "name": "John Doe"
  }
}
```

### 3. Refresh Access Token

**Endpoint:** `POST /api/v1/auth/refresh`

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response:** `200 OK`
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 3600
}
```

### 4. Logout

**Endpoint:** `POST /api/v1/auth/logout`

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response:** `200 OK`
```json
{
  "message": "logged out successfully"
}
```

### 5. Get Current User (Protected)

**Endpoint:** `GET /api/v1/auth/me`

**Headers:**
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Response:** `200 OK`
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "john@example.com",
  "name": "John Doe"
}
```

## Token Details

### Access Token
- **Type:** JWT (JSON Web Token)
- **Expiration:** 1 hour (configurable via `JWT_EXPIRY` env var)
- **Use:** Include in Authorization header for protected routes
- **Format:** `Authorization: Bearer <access_token>`

### Refresh Token
- **Type:** JWT (JSON Web Token)
- **Expiration:** 7 days (168 hours)
- **Use:** Obtain new access tokens without re-authentication
- **Storage:** Stored in database sessions table

## Environment Variables

```bash
# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
JWT_EXPIRY=24  # Access token expiration in hours

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=mytodo
DB_PASSWORD=mytodo_dev_password
DB_NAME=mytodo_dev
```

## Database Schema

### Users Table
```sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_login_at TIMESTAMP,
    is_active BOOLEAN NOT NULL DEFAULT true
);
```

### Sessions Table
```sql
CREATE TABLE sessions (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    refresh_token TEXT NOT NULL,
    user_agent TEXT,
    ip_address VARCHAR(45),
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    revoked_at TIMESTAMP
);
```

## Security Features

1. **Password Hashing:** All passwords are hashed using bcrypt with a cost factor of 10
2. **Token Signing:** JWT tokens are signed using HMAC-SHA256
3. **Session Tracking:** Each refresh token is linked to a session with device info
4. **Token Expiration:** Both access and refresh tokens have expiration times
5. **Session Revocation:** Logout revokes the session, invalidating the refresh token
6. **Protected Routes:** Middleware validates JWT tokens before allowing access

## Error Responses

All endpoints return appropriate HTTP status codes and error messages:

- `400 Bad Request` - Invalid request body or validation errors
- `401 Unauthorized` - Invalid credentials or expired/invalid tokens
- `404 Not Found` - User not found
- `500 Internal Server Error` - Server-side errors

**Error Response Format:**
```json
{
  "error": "detailed error message"
}
```

## Testing with cURL

### Register
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","password":"SecurePassword123"}'
```

### Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"SecurePassword123"}'
```

### Get Current User
```bash
curl -X GET http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

### Refresh Token
```bash
curl -X POST http://localhost:8080/api/v1/auth/refresh \
  -H "Content-Type: application/json" \
  -d '{"refresh_token":"YOUR_REFRESH_TOKEN"}'
```

### Logout
```bash
curl -X POST http://localhost:8080/api/v1/auth/logout \
  -H "Content-Type: application/json" \
  -d '{"refresh_token":"YOUR_REFRESH_TOKEN"}'
```

## Setup Instructions

1. **Install Dependencies:**
   ```bash
   cd d:/projects/mytodo
   go mod tidy
   ```

2. **Start Database:**
   ```bash
   docker-compose up -d postgres
   ```

3. **Initialize Database:**
   The database schema is automatically created when the PostgreSQL container starts for the first time.

4. **Start the API Server:**
   ```bash
   cd apps/api
   go run cmd/server/main.go
   ```

5. **Verify Server is Running:**
   ```bash
   curl http://localhost:8080/health
   ```

## Architecture

The authentication service follows Clean Architecture principles with clear layer separation:

- **Domain Layer:** Entities, repositories interfaces, and business logic services
- **Application Layer:** Use cases and command handlers
- **Infrastructure Layer:** Repository implementations, database access
- **Interface Layer:** HTTP controllers, DTOs, middleware

### Key Components

1. **Entities:**
   - `User`: User account entity
   - `Session`: Refresh token session entity

2. **Services:**
   - `AuthService`: Core authentication business logic
   - `JWTService`: JWT token generation and validation
   - `PasswordService`: Password hashing and verification

3. **Repositories:**
   - `UserRepository`: User data persistence
   - `SessionRepository`: Session data persistence

4. **Controllers:**
   - `AuthController`: HTTP endpoints for authentication

5. **Middleware:**
   - `AuthMiddleware`: JWT token validation for protected routes

## Best Practices

1. **Always use HTTPS in production** to protect tokens in transit
2. **Store refresh tokens securely** on the client side (httpOnly cookies recommended)
3. **Rotate refresh tokens** periodically for enhanced security
4. **Implement rate limiting** on login endpoint to prevent brute force attacks
5. **Log authentication events** for security auditing
6. **Use strong JWT secrets** in production (32+ characters)
7. **Implement token blacklisting** for enhanced security (optional)

## Future Enhancements

- Email verification during registration
- Password reset via email
- Two-factor authentication (2FA)
- OAuth integration (Google, GitHub)
- Rate limiting on authentication endpoints
- Account lockout after failed login attempts
- Password complexity requirements
- Token blacklisting for instant revocation
