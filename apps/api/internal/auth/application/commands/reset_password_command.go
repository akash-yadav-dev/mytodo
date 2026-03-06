// Package commands implements Command pattern for write operations.
//
// This file defines the password reset command.

package commands

// ResetPasswordCommand represents a password reset request.
// In production applications, reset password commands typically include:
// - Reset token (from email link)
// - New password
// - Password confirmation
// - Optional: user ID (derived from token)
// - Optional: invalidate all sessions flag
// - IP address for security logging
// - User-Agent for audit trail
