package persistence_test

import (
	"context"
	"testing"
	"time"

	"mytodo/apps/api/internal/auth/domain/entity"
	"mytodo/apps/api/internal/auth/infrastructure/persistence"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func TestSessionRepository_CreateAndFindByRefreshToken(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	repo := persistence.NewSessionRepository(db)

	sessionID := uuid.New()
	userID := uuid.New()
	now := time.Now()
	session := &entity.Session{
		ID:           sessionID,
		UserID:       userID,
		RefreshToken: "refresh-token",
		UserAgent:    "agent",
		IPAddress:    "127.0.0.1",
		ExpiresAt:    now.Add(24 * time.Hour),
		CreatedAt:    now,
	}

	mock.ExpectExec("INSERT INTO sessions").
		WithArgs(session.ID, session.UserID, session.RefreshToken, session.UserAgent, session.IPAddress, session.ExpiresAt, session.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := repo.Create(context.Background(), session); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	rows := sqlmock.NewRows([]string{
		"id", "user_id", "refresh_token", "user_agent", "ip_address", "expires_at", "created_at", "revoked_at",
	}).AddRow(session.ID, session.UserID, session.RefreshToken, session.UserAgent, session.IPAddress, session.ExpiresAt, session.CreatedAt, nil)

	mock.ExpectQuery("SELECT id, user_id, refresh_token").WithArgs(session.RefreshToken).WillReturnRows(rows)

	found, err := repo.FindByRefreshToken(context.Background(), session.RefreshToken)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if found.RefreshToken != session.RefreshToken {
		t.Fatalf("expected refresh token to match")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expected all SQL expectations to be met: %v", err)
	}
}

func TestSessionRepository_UpdateRevokedAt(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	repo := persistence.NewSessionRepository(db)

	sessionID := uuid.New()
	revokedAt := time.Now()

	mock.ExpectExec("UPDATE sessions").WithArgs(&revokedAt, sessionID).WillReturnResult(sqlmock.NewResult(1, 1))

	session := &entity.Session{ID: sessionID, RevokedAt: &revokedAt}

	if err := repo.Update(context.Background(), session); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expected all SQL expectations to be met: %v", err)
	}
}
