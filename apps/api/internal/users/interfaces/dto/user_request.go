// Package dto defines Data Transfer Objects for users module.

package dto

// Request DTOs for user operations.
//
// Example structures:
//   type CreateUserRequest struct {
//       Email       string `json:"email" validate:"required,email"`
//       Username    string `json:"username" validate:"required,alphanum,min=3,max=30"`
//       DisplayName string `json:"display_name" validate:"required"`
//   }
//
//   type UpdateUserRequest struct {
//       DisplayName string `json:"display_name,omitempty"`
//       Avatar      string `json:"avatar,omitempty" validate:"omitempty,url"`
//       Bio         string `json:"bio,omitempty" validate:"omitempty,max=500"`
//   }
