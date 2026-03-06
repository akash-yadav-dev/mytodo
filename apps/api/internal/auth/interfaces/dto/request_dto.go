// Package dto defines Data Transfer Objects for the auth module.
//
// This file contains request DTOs for various auth operations.

package dto

// Request DTOs represent incoming API requests.
// In production applications, request DTOs typically:
// - Define all required and optional fields
// - Include validation tags (required, email, min, max)
// - Provide binding rules for frameworks (JSON, form data)
// - Document field constraints
// - Include example values for API docs
// - Handle multiple content types (JSON, form-urlencoded)
// Common request DTOs:
// - LoginRequest, RegisterRequest, UpdateUserRequest
// - ForgotPasswordRequest, ResetPasswordRequest
// - VerifyEmailRequest, ResendVerificationRequest
