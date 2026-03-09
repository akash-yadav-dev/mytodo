// Package entity defines core domain entities for the users module.
//
// Domain entities represent user-related business objects in Clean Architecture.
// These are rich domain models containing business logic and invariants.
//
// In production-grade applications, this folder typically contains:
// - User aggregate root with identity and core attributes
// - Value objects (Email, Username, Avatar)
// - Business rules and validation logic
// - Entity lifecycle methods
// - Domain events related to user changes

package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// User represents a user profile in the system (extends auth.User).
// This is the domain model for user profile and management.
type User struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`  // Foreign key to auth.users
	Username    *string   `json:"username"` // Unique username (optional)
	DisplayName string    `json:"display_name"`
	AvatarURL   string    `json:"avatar_url"`
	Bio         string    `json:"bio"`
	Location    string    `json:"location"`
	Website     string    `json:"website"`
	Phone       string    `json:"phone"`
	Timezone    string    `json:"timezone"`
	Language    string    `json:"language"`
	Theme       string    `json:"theme"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewUserProfile creates a new user profile
func NewUserProfile(userID uuid.UUID, displayName string) *User {
	now := time.Now()
	return &User{
		ID:          uuid.New(),
		UserID:      userID,
		DisplayName: displayName,
		Timezone:    "UTC",
		Language:    "en",
		Theme:       "light",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// UpdateProfile updates user profile information
func (u *User) UpdateProfile(displayName, bio, location, website string) error {
	if displayName == "" {
		return errors.New("display name cannot be empty")
	}
	u.DisplayName = displayName
	u.Bio = bio
	u.Location = location
	u.Website = website
	u.UpdatedAt = time.Now()
	return nil
}

// ChangeAvatar updates the user's avatar URL
func (u *User) ChangeAvatar(avatarURL string) error {
	if avatarURL == "" {
		return errors.New("avatar URL cannot be empty")
	}
	u.AvatarURL = avatarURL
	u.UpdatedAt = time.Now()
	return nil
}

// SetUsername sets or updates the username
func (u *User) SetUsername(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}
	if len(username) < 3 || len(username) > 30 {
		return errors.New("username must be between 3 and 30 characters")
	}
	u.Username = &username
	u.UpdatedAt = time.Now()
	return nil
}

// UpdatePreferences updates user preferences
func (u *User) UpdatePreferences(timezone, language, theme string) error {
	if timezone != "" {
		u.Timezone = timezone
	}
	if language != "" {
		u.Language = language
	}
	if theme != "" {
		u.Theme = theme
	}
	u.UpdatedAt = time.Now()
	return nil
}
