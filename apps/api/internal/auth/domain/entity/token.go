// Package entity defines the core domain entities for the auth module.
//
// This file contains token-related entities that represent authentication tokens
// used for session management and API access control.

package entity

// Token represents an authentication token (JWT, refresh token, or API token).
// In production applications, token entities typically include:
// - Token value (hashed for security)
// - Token type (access_token, refresh_token, api_key)
// - Associated user ID
// - Expiration timestamp
// - Scope/permissions granted by this token
// - Token status (active, revoked, expired)
// - Device/client information (for device tracking)
// - IP address and user agent (for security auditing)
// - Created and revoked timestamps
// - Token family/chain (for refresh token rotation)
