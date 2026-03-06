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

// User represents a user in the system (different from auth.User).
// This is the domain model for user profile and management.
//
// In production applications, User entities typically include:
// - ID: unique identifier (UUID)
// - Email: unique email address
// - Username: unique username
// - DisplayName: user's display name
// - Avatar: profile picture URL
// - Bio: user biography/description
// - Status: active, suspended, deleted
// - Metadata: custom fields, preferences
// - Timestamps: created_at, updated_at
//
// Example structure:
//   type User struct {
//       ID          string    `json:"id"`
//       Email       string    `json:"email"`
//       Username    string    `json:"username"`
//       DisplayName string    `json:"display_name"`
//       Avatar      string    `json:"avatar"`
//       Bio         string    `json:"bio"`
//       Status      string    `json:"status"`
//       CreatedAt   time.Time `json:"created_at"`
//       UpdatedAt   time.Time `json:"updated_at"`
//   }
//
// Example methods:
//   func (u *User) UpdateProfile(name, bio string) error
//   func (u *User) ChangeAvatar(avatarURL string) error
//   func (u *User) Deactivate() error
//   func (u *User) IsActive() bool
