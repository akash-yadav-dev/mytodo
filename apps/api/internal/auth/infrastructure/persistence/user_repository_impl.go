// Package persistence provides concrete implementations of domain repository interfaces.
//
// Infrastructure/Persistence layer in Clean Architecture:
// - Implements repository interfaces defined in the domain layer
// - Contains database-specific code (SQL, ORM, NoSQL queries)
// - Handles data mapping between database models and domain entities
// - Manages database connections and transactions
//
// In production applications, this folder typically contains:
// - Repository implementations using specific databases (PostgreSQL, MySQL, MongoDB)
// - ORM model definitions (GORM, sqlx, or raw SQL)
// - Data mappers (convert DB models to domain entities and vice versa)
// - Database query builders
// - Migration-related utility functions
// - Connection pooling configuration
//
// Best practices:
// - Isolate database technology from domain layer
// - Handle database errors and convert to domain errors
// - Implement efficient queries (indexes, joins)
// - Use prepared statements to prevent SQL injection
// - Handle NULL values appropriately
// - Implement proper transaction management

package persistence

import (
	"context"
	"database/sql"
	"errors"
	"mytodo/apps/api/internal/auth/domain/entity"
	"mytodo/apps/api/internal/auth/domain/repository"

	"github.com/google/uuid"
)

// UserRepositoryImpl is the PostgreSQL implementation of UserRepository interface.
// In production applications, repository implementations typically:
// - Use connection pooling for performance
// - Implement all interface methods defined in domain/repository
// - Map between database rows/documents and domain entities
// - Handle database-specific errors (unique constraints, foreign keys)
// - Implement optimized queries with proper indexes
// - Cache frequently accessed data
// - Use transactions for multi-step operations
// - Log slow queries for performance monitoring

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	query := `
		INSERT INTO users (id, email, password_hash, name, created_at, updated_at, is_active)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.Name,
		user.CreatedAt,
		user.UpdatedAt,
		user.IsActive,
	)
	return err
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	query := `
		SELECT id, email, password_hash, name, created_at, updated_at, last_login_at, is_active
		FROM users
		WHERE id = $1
	`
	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLoginAt,
		&user.IsActive,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT id, email, password_hash, name, created_at, updated_at, last_login_at, is_active
		FROM users
		WHERE email = $1
	`
	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLoginAt,
		&user.IsActive,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
	query := `
		UPDATE users
		SET name = $1, updated_at = $2, last_login_at = $3
		WHERE id = $4
	`
	_, err := r.db.ExecContext(ctx, query,
		user.Name,
		user.UpdatedAt,
		user.LastLoginAt,
		user.ID,
	)
	return err
}

func (r *UserRepositoryImpl) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)
	return exists, err
}
