// Package persistence provides concrete implementations of repository interfaces.
//
// This file implements session repository using PostgreSQL/Redis.

package persistence

import (
	"context"
	"database/sql"
	"errors"
	"mytodo/apps/api/internal/auth/domain/entity"
	"mytodo/apps/api/internal/auth/domain/repository"

	"github.com/google/uuid"
)

// SessionRepositoryImpl implements SessionRepository interface.
// In production applications, session repositories typically:
// - Use Redis or in-memory cache for fast session lookups
// - Implement TTL (time-to-live) for automatic expiration
// - Fall back to database for persistent sessions
// - Handle concurrent session access (locking if needed)
// - Implement efficient cleanup of expired sessions
// - Support both short-lived and extended sessions
// - Serialize session data efficiently (JSON, MessagePack)

type SessionRepositoryImpl struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) repository.SessionRepository {
	return &SessionRepositoryImpl{db: db}
}

func (r *SessionRepositoryImpl) Create(ctx context.Context, session *entity.Session) error {
	query := `
		INSERT INTO sessions (id, user_id, refresh_token, user_agent, ip_address, expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.ExecContext(ctx, query,
		session.ID,
		session.UserID,
		session.RefreshToken,
		session.UserAgent,
		session.IPAddress,
		session.ExpiresAt,
		session.CreatedAt,
	)
	return err
}

func (r *SessionRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Session, error) {
	query := `
		SELECT id, user_id, refresh_token, user_agent, ip_address, expires_at, created_at, revoked_at
		FROM sessions
		WHERE id = $1
	`
	session := &entity.Session{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&session.ID,
		&session.UserID,
		&session.RefreshToken,
		&session.UserAgent,
		&session.IPAddress,
		&session.ExpiresAt,
		&session.CreatedAt,
		&session.RevokedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("session not found")
	}
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (r *SessionRepositoryImpl) FindByRefreshToken(ctx context.Context, token string) (*entity.Session, error) {
	query := `
		SELECT id, user_id, refresh_token, user_agent, ip_address, expires_at, created_at, revoked_at
		FROM sessions
		WHERE refresh_token = $1
	`
	session := &entity.Session{}
	err := r.db.QueryRowContext(ctx, query, token).Scan(
		&session.ID,
		&session.UserID,
		&session.RefreshToken,
		&session.UserAgent,
		&session.IPAddress,
		&session.ExpiresAt,
		&session.CreatedAt,
		&session.RevokedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("session not found")
	}
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (r *SessionRepositoryImpl) Update(ctx context.Context, session *entity.Session) error {
	query := `
		UPDATE sessions
		SET revoked_at = $1
		WHERE id = $2
	`
	_, err := r.db.ExecContext(ctx, query, session.RevokedAt, session.ID)
	return err
}

func (r *SessionRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM sessions WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *SessionRepositoryImpl) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	query := `DELETE FROM sessions WHERE user_id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}
