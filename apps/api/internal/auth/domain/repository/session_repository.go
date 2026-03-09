// Package repository defines data access interfaces for the auth domain.
//
// This file defines the contract for session data persistence.

package repository

import (
	"context"
	"mytodo/apps/api/internal/auth/domain/entity"

	"github.com/google/uuid"
)

// SessionRepository defines data access methods for session entities.
// In production applications, session repositories typically provide:
// - CreateSession(session) - persist new session
// - FindByID(sessionID) - retrieve session by ID
// - FindByUserID(userID) - get all sessions for a user
// - FindActiveByUserID(userID) - get only active sessions
// - UpdateActivity(sessionID, timestamp) - update last activity
// - InvalidateSession(sessionID) - mark session as logged out
// - InvalidateAllUserSessions(userID) - force logout all user sessions
// - DeleteExpiredSessions() - cleanup expired sessions
// - CountActiveByUserID(userID) - count concurrent sessions
// - FindByToken(token) - lookup session by token
// - UpdateMetadata(sessionID, metadata) - update session data

type SessionRepository interface {
	Create(ctx context.Context, session *entity.Session) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Session, error)
	FindByRefreshToken(ctx context.Context, token string) (*entity.Session, error)
	Update(ctx context.Context, session *entity.Session) error
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteByUserID(ctx context.Context, userID uuid.UUID) error
}
