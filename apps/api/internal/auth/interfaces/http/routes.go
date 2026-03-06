// Package http provides HTTP/REST API endpoints for the auth module.
//
// This file defines route registration and grouping.

package http

// Routes file registers HTTP routes for the auth module.
// In production applications, route files typically:
// - Define route groups (/api/v1/auth, /api/v1/users)
// - Register all endpoints with their handlers
// - Apply middleware to routes (authentication, authorization)
// - Define route-specific middleware (rate limiting on login)
// - Configure CORS for specific routes
// - Set up route-level metrics and logging
// - Define API versioning strategy
// - Document routes (for API documentation generation)
