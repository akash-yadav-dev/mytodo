// Package cache provides caching implementations for the auth module.
//
// This file implements token caching for blacklisting and validation.

package cache

// TokenCache manages token-related caching operations.
// In production applications, token caches typically:
// - Implement token blacklist (revoked tokens)
// - Cache token validation results to reduce computation
// - Store refresh token families for rotation tracking
// - Implement token rate limiting data
// - Cache public keys for JWT verification
// - Handle distributed token state across servers
// - Set TTL matching token expiration
// - Support atomic operations for token refresh
