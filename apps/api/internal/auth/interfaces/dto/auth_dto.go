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

// Auth DTOs for authentication operations.
// In production applications, auth DTOs typically include:
// - LoginRequest: email, password
// - LoginResponse: access_token, refresh_token, expires_in, user
// - RegisterRequest: email, password, name, etc.
// - RegisterResponse: user, tokens
// - RefreshTokenRequest: refresh_token
// - RefreshTokenResponse: access_token, expires_in
// - ResetPasswordRequest: token, new_password
