// Package entity defines the core domain entities for the auth module.
//
// This file contains session-related entities for managing user sessions
// and tracking active authenticated connections.

package entity

// Session represents an active user session in the system.
// In production applications, session entities typically include:
// - Unique session ID
// - Associated user ID
// - Session start and expiry times
// - Last activity timestamp (for idle timeout)
// - Device fingerprint (browser, OS, device type)
// - IP address (for security monitoring)
// - Geolocation data (country, city)
// - Session metadata (login method, MFA status)
// - Session status (active, expired, logged_out, force_killed)
// - Concurrent session limits and handling
// - Remember me flag (for extended sessions)
