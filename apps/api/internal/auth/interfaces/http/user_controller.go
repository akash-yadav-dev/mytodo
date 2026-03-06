// Package http provides HTTP/REST API endpoints for the auth module.
//
// This file implements user management HTTP endpoints.

package http

// UserController handles HTTP endpoints for user management.
// In production applications, user controllers typically implement:
// - GET /users - list users with pagination and filters
// - GET /users/:id - get user by ID
// - PUT /users/:id - update user
// - DELETE /users/:id - delete/deactivate user
// - GET /users/:id/permissions - get user permissions
// - POST /users/:id/roles - assign role to user
// - DELETE /users/:id/roles/:roleId - remove role from user
// - GET /users/search - search users
// - PATCH /users/:id/password - change password
