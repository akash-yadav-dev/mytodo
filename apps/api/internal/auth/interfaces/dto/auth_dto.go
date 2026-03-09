// Package dto defines Data Transfer Objects for the auth module's interface layer.
//
// DTOs (Data Transfer Objects) in Clean Architecture:
// - Transfer data between layers without exposing domain entities
// - Define the contract for API requests and responses
// - Provide data validation and transformation
// - Decouple external API from internal domain model
//
// In production applications, this folder typically contains:
// - Request DTOs (input validation, deserialization)
// - Response DTOs (output formatting, serialization)
// - Validation tags (required, email, min/max length)
// - JSON/XML field mappings
// - API documentation annotations
// - Data transformation methods (to/from domain entities)
//
// Best practices:
// - Never expose domain entities directly through APIs
// - Use separate DTOs for request and response (different fields)
// - Add validation tags for input validation
// - Omit sensitive fields in responses (passwords, tokens)
// - Use clear, consistent naming conventions
// - Version DTOs when API changes
// - Document fields for API docs generation

package dto

import "mytodo/apps/api/internal/auth/domain/entity"

// LoginRequest represents login credentials
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterRequest represents registration data
type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=2"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RefreshTokenRequest represents refresh token request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// AuthResponse represents authentication response with tokens
type AuthResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	TokenType    string   `json:"token_type"`
	ExpiresIn    int64    `json:"expires_in"`
	User         *UserDTO `json:"user"`
}

// UserDTO represents user information in responses
type UserDTO struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// ToUserDTO converts domain User entity to DTO
func ToUserDTO(user *entity.User) *UserDTO {
	return &UserDTO{
		ID:    user.ID.String(),
		Email: user.Email,
		Name:  user.Name,
	}
}

// MessageResponse represents a simple message response
type MessageResponse struct {
	Message string `json:"message"`
}
