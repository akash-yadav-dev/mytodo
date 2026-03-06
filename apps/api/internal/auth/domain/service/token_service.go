// Package service contains domain services for the auth module.
//
// This file implements token management business logic.

package service

// TokenService handles token generation, validation, and lifecycle management.
// In production applications, token services typically implement:
// - JWT token generation with claims and signing
// - Token validation and signature verification
// - Token refresh logic (refresh token rotation)
// - Token revocation and blacklisting
// - Token expiration handling
// - Custom claims management
// - Token scope and permission encoding
// - API key generation and validation
// - Token family tracking (for refresh token chains)
// - Cryptographic operations for token security
