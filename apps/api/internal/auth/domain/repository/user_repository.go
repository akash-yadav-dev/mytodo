// Package repository defines interfaces (contracts) for data access in the auth domain.
//
// Repository interfaces in Clean Architecture:
// - Define the contract for data access without implementation details
// - Allow the domain layer to remain independent of infrastructure
// - Enable dependency inversion (domain depends on abstractions, not concrete implementations)
// - Support testing through mocking/stubbing
//
// In production applications, repository interfaces typically define methods for:
// - CRUD operations (Create, Read, Update, Delete)
// - Query methods (Find by ID, email, filters)
// - Bulk operations (batch inserts, updates)
// - Existence checks
// - Complex queries specific to business needs
// - Transaction boundaries (when needed)
//
// Note: The actual implementation lives in infrastructure/persistence,
// which depends on specific database technology (PostgreSQL, MongoDB, etc.)

package repository

import (
	"context"
	"mytodo/apps/api/internal/auth/domain/entity"

	"github.com/google/uuid"
)

// UserRepository defines data access methods for user entities.
// In production applications, user repositories typically provide:
// - FindByID(id) - retrieve user by unique identifier
// - FindByEmail(email) - lookup user by email (for login)
// - FindByUsername(username) - lookup by username
// - Create(user) - persist new user
// - Update(user) - update existing user
// - Delete(id) - soft or hard delete
// - List(filters, pagination) - query users with filters
// - ExistsByEmail(email) - check email uniqueness
// - UpdatePassword(id, hashedPassword) - specialized password update
// - UpdateLastLogin(id, timestamp) - track login activity
// - ActivateAccount(id) - account activation
// - Search(query) - full-text search across user fields

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}
