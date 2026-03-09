package persistence_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"mytodo/apps/api/internal/auth/domain/entity"
	"mytodo/apps/api/internal/auth/infrastructure/persistence"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func TestUserRepository_CreateAndFindByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	repo := persistence.NewUserRepository(db)

	userID := uuid.New()
	now := time.Now()
	user := &entity.User{
		ID:           userID,
		Email:        "user@example.com",
		PasswordHash: "hashed",
		Name:         "User",
		CreatedAt:    now,
		UpdatedAt:    now,
		IsActive:     true,
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.ID, user.Email, user.PasswordHash, user.Name, user.CreatedAt, user.UpdatedAt, user.IsActive).
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := repo.Create(context.Background(), user); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	rows := sqlmock.NewRows([]string{
		"id", "email", "password_hash", "name", "created_at", "updated_at", "last_login_at", "is_active",
	}).AddRow(user.ID, user.Email, user.PasswordHash, user.Name, user.CreatedAt, user.UpdatedAt, sql.NullTime{}, user.IsActive)

	mock.ExpectQuery("SELECT id, email, password_hash").WithArgs(user.Email).WillReturnRows(rows)

	found, err := repo.FindByEmail(context.Background(), user.Email)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if found.Email != user.Email {
		t.Fatalf("expected email to match")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expected all SQL expectations to be met: %v", err)
	}
}
