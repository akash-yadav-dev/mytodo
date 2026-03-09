// Package repository defines data access interfaces for the users domain.
//
// Repository interfaces define contracts for data persistence without
// specifying implementation details (database-agnostic).

package repository

import (
	"context"
	"mytodo/apps/api/internal/users/domain/entity"

	"github.com/google/uuid"
)

// UserRepository defines data access methods for user profile entities.
type UserRepository interface {
	// Profile operations
	CreateProfile(ctx context.Context, user *entity.User) error
	FindProfileByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	FindProfileByUserID(ctx context.Context, userID uuid.UUID) (*entity.User, error)
	FindProfileByUsername(ctx context.Context, username string) (*entity.User, error)
	UpdateProfile(ctx context.Context, user *entity.User) error
	DeleteProfile(ctx context.Context, userID uuid.UUID) error
	ListProfiles(ctx context.Context, page, limit int) ([]*entity.User, int, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
	SearchProfiles(ctx context.Context, query string, limit int) ([]*entity.User, error)

	// Preference operations
	CreatePreferences(ctx context.Context, pref *entity.Preference) error
	FindPreferencesByUserID(ctx context.Context, userID uuid.UUID) (*entity.Preference, error)
	UpdatePreferences(ctx context.Context, pref *entity.Preference) error
}
