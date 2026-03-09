// Package repository defines data access interfaces for the users domain.
//
// This file defines the contract for profile data persistence.

package repository

import "mytodo/apps/api/internal/users/domain/entity"

// ProfileRepository defines data access methods for profile entities.
//
// In production applications, profile repositories typically provide:
// - Find profile by user ID
// - Create new profile (typically on user registration)
// - Update profile information
// - Delete profile (cascade on user deletion)
//
// Example interface:
type UserProfileRepository interface {
	Create(profile *entity.Profile) error
	FindByUserID(userID string) (*entity.Profile, error)
	Update(profile *entity.Profile) error
	Delete(userID string) error
}

//
// Example usage:
//   profile, err := repo.FindByUserID("user-123")
//   // Returns: &Profile{UserID: "user-123", FirstName: "John", ...}, nil
//   // Returns: nil, ErrProfileNotFound if profile doesn't exist
//
//   err := repo.Update(&Profile{
//       UserID:    "user-123",
//       FirstName: "Jane",
//       LastName:  "Doe",
//   })
//   // Returns: nil on success, error on failure
