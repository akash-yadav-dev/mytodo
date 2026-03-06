// Package http provides HTTP/REST API endpoints for the auth module.
//
// This file implements HTTP middleware for the auth module.

package http

// Middlewares file implements HTTP middleware functions.
// In production applications, auth middleware typically includes:
// - JWT authentication middleware (verify and extract token)
// - Authorization middleware (check permissions/roles)
// - Rate limiting (prevent brute force attacks)
// - Request logging and tracing
// - CORS handling
// - Request ID injection
// - Timeout middleware
// - Security headers (HSTS, CSP, X-Frame-Options)
// - Request size limits
// - API key validation
