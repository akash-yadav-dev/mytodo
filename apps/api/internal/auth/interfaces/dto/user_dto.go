// Package dto defines Data Transfer Objects for the auth module.
//
// This file contains user-related DTOs.

package dto

// User DTOs for user management operations.
// In production applications, user DTOs typically include:
// - UserResponse: id, email, name, avatar, roles, created_at (no password!)
// - UserListResponse: users[], total, page, per_page
// - UpdateUserRequest: name, avatar, bio
// - ChangePasswordRequest: old_password, new_password
// - UserPermissionsResponse: permissions[], roles[]
// - Validation tags: `json:"email" validate:"required,email"`
