// Package dto defines Data Transfer Objects for the auth module.
//
// This file contains response DTOs for API responses.

package dto

// Response DTOs represent outgoing API responses.
// In production applications, response DTOs typically:
// - Define the exact structure of API responses
// - Omit sensitive fields (password hashes, internal IDs)
// - Use consistent naming conventions (snake_case or camelCase)
// - Include metadata (timestamps, pagination info)
// - Format dates/times consistently (ISO 8601)
// - Handle nested objects and collections
// - Support multiple serialization formats (JSON, XML)
// Common response DTOs:
// - AuthResponse, UserResponse, ErrorResponse
// - PaginatedResponse, SuccessResponse
// - Standard error format: {error: string, code: string, details: {}}
