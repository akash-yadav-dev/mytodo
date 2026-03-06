// Package http provides HTTP/REST API endpoints for the auth module.
//
// Interfaces/HTTP layer in Clean Architecture:
// - Handles HTTP requests and responses
// - Implements REST API endpoints
// - Manages request/response serialization (JSON, XML)
// - Applies HTTP-specific concerns (routing, middleware)
//
// In production applications, this folder typically contains:
// - HTTP controllers/handlers for each resource
// - Route definitions and registration
// - HTTP middleware (auth, logging, CORS, rate limiting)
// - Request validation and binding
// - Response formatting and error handling
// - API versioning logic
// - OpenAPI/Swagger documentation
//
// Best practices:
// - Keep controllers thin - delegate to application handlers
// - Use DTOs for request/response (don't expose domain entities)
// - Handle HTTP status codes appropriately
// - Implement proper error responses (RFC 7807)
// - Apply security headers
// - Implement request ID tracking
// - Use content negotiation

package http

// AuthController handles HTTP endpoints for authentication operations.
// In production applications, auth controllers typically implement:
// - POST /auth/register - user registration
// - POST /auth/login - user authentication
// - POST /auth/logout - session termination
// - POST /auth/refresh - token refresh
// - POST /auth/forgot-password - password reset request
// - POST /auth/reset-password - password reset completion
// - GET /auth/me - current user info
// - POST /auth/verify-email - email verification
// - OAuth endpoints (/auth/google, /auth/github, /auth/callback)
