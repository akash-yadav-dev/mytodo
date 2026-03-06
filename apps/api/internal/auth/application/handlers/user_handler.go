// Package handlers contains application-layer handlers for use case orchestration.
//
// This file implements handlers for user management operations.

package handlers

// UserHandler orchestrates user management use cases.
// In production applications, user handlers typically implement:
// - Handle(GetUserQuery) - retrieve user details
// - Handle(ListUsersQuery) - list with filters and pagination
// - Handle(UpdateUserCommand) - update user profile
// - Handle(DeleteUserCommand) - deactivate/delete user
// - Handle(GetPermissionsQuery) - retrieve user permissions
// - Coordinate with user service and repositories
// - Apply authorization checks
// - Manage data transformations (entity to DTO)
// - Handle validation and business rule enforcement
