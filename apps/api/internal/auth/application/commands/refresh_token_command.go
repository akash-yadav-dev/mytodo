// Package commands implements Command pattern for write operations.
//
// This file defines the token refresh command.

package commands

// RefreshTokenCommand represents a request to refresh an access token.
// In production applications, refresh token commands typically include:
// - Refresh token value
// - Optional: access token (for token binding)
// - Optional: device fingerprint (for security)
// - IP address for security validation
// - User-Agent for device tracking
