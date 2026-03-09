// Package entity defines the core domain entities for the auth module.
//
// This file contains session-related entities for managing user sessions
// and tracking active authenticated connections.

package entity

import (
	"time"

	"github.com/google/uuid"
)

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

type Session struct {
	ID           uuid.UUID  `json:"id"`
	UserID       uuid.UUID  `json:"user_id"`
	RefreshToken string     `json:"refresh_token"`
	UserAgent    string     `json:"user_agent"`
	IPAddress    string     `json:"ip_address"`
	ExpiresAt    time.Time  `json:"expires_at"`
	CreatedAt    time.Time  `json:"created_at"`
	RevokedAt    *time.Time `json:"revoked_at,omitempty"`
}

// NewSession creates a new session for a user
func NewSession(userID uuid.UUID, refreshToken, userAgent, ipAddress string, expiresAt time.Time) *Session {
	return &Session{
		ID:           uuid.New(),
		UserID:       userID,
		RefreshToken: refreshToken,
		UserAgent:    userAgent,
		IPAddress:    ipAddress,
		ExpiresAt:    expiresAt,
		CreatedAt:    time.Now(),
	}
}

// IsExpired checks if the session has expired
func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}

// IsRevoked checks if the session has been revoked
func (s *Session) IsRevoked() bool {
	return s.RevokedAt != nil
}

// Revoke marks the session as revoked
func (s *Session) Revoke() {
	now := time.Now()
	s.RevokedAt = &now
}
