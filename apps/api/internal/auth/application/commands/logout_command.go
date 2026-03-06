// Package commands implements Command pattern for write operations.
//
// This file defines the logout command.

package commands

// LogoutCommand represents a user logout request.
// In production applications, logout commands typically include:
// - Session ID or token to invalidate
// - Optional: User ID (for validation)
// - Optional: logout type (single session vs all sessions)
// - Optional: device ID (for device-specific logout)
// - Timestamp for audit logging
