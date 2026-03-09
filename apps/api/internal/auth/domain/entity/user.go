package entity

import (
	"time"

	"github.com/google/uuid"
)

// User represents an authenticated user in the system.
// In production applications, user entities typically include:
// - Unique identifier (ID, UUID)
// - Authentication credentials (email, hashed password)
// - Profile information (name, avatar)
// - Account status (active, suspended, deleted)
// - Security metadata (failed login attempts, MFA settings)
// - Timestamps (created_at, updated_at, last_login)
// - Relationships to roles and permissions

type User struct {
	ID           uuid.UUID  `json:"id"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"` // Never expose password hash in JSON
	Name         string     `json:"name"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	LastLoginAt  *time.Time `json:"last_login_at,omitempty"`
	IsActive     bool       `json:"is_active"`
}

// NewUser creates a new user with generated ID and timestamps
func NewUser(email, name, passwordHash string) *User {
	now := time.Now()
	return &User{
		ID:           uuid.New(),
		Email:        email,
		Name:         name,
		PasswordHash: passwordHash,
		CreatedAt:    now,
		UpdatedAt:    now,
		IsActive:     true,
	}
}

// UpdateLastLogin updates the last login timestamp
func (u *User) UpdateLastLogin() {
	now := time.Now()
	u.LastLoginAt = &now
	u.UpdatedAt = now
}
