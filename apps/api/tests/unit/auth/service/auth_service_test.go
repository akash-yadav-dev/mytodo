package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"mytodo/apps/api/internal/auth/domain/entity"
	"mytodo/apps/api/internal/auth/domain/service"
	"mytodo/apps/api/pkg/security"

	"github.com/google/uuid"
)

type fakeUserRepo struct {
	usersByID    map[uuid.UUID]*entity.User
	usersByEmail map[string]*entity.User
}

func newFakeUserRepo() *fakeUserRepo {
	return &fakeUserRepo{
		usersByID:    map[uuid.UUID]*entity.User{},
		usersByEmail: map[string]*entity.User{},
	}
}

func (r *fakeUserRepo) Create(ctx context.Context, user *entity.User) error {
	r.usersByID[user.ID] = user
	r.usersByEmail[user.Email] = user
	return nil
}

func (r *fakeUserRepo) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, ok := r.usersByID[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *fakeUserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, ok := r.usersByEmail[email]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *fakeUserRepo) Update(ctx context.Context, user *entity.User) error {
	r.usersByID[user.ID] = user
	r.usersByEmail[user.Email] = user
	return nil
}

func (r *fakeUserRepo) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	_, ok := r.usersByEmail[email]
	return ok, nil
}

type fakeSessionRepo struct {
	sessionsByID    map[uuid.UUID]*entity.Session
	sessionsByToken map[string]*entity.Session
}

func newFakeSessionRepo() *fakeSessionRepo {
	return &fakeSessionRepo{
		sessionsByID:    map[uuid.UUID]*entity.Session{},
		sessionsByToken: map[string]*entity.Session{},
	}
}

func (r *fakeSessionRepo) Create(ctx context.Context, session *entity.Session) error {
	r.sessionsByID[session.ID] = session
	r.sessionsByToken[session.RefreshToken] = session
	return nil
}

func (r *fakeSessionRepo) FindByID(ctx context.Context, id uuid.UUID) (*entity.Session, error) {
	session, ok := r.sessionsByID[id]
	if !ok {
		return nil, errors.New("session not found")
	}
	return session, nil
}

func (r *fakeSessionRepo) FindByRefreshToken(ctx context.Context, token string) (*entity.Session, error) {
	session, ok := r.sessionsByToken[token]
	if !ok {
		return nil, errors.New("session not found")
	}
	return session, nil
}

func (r *fakeSessionRepo) Update(ctx context.Context, session *entity.Session) error {
	r.sessionsByID[session.ID] = session
	r.sessionsByToken[session.RefreshToken] = session
	return nil
}

func (r *fakeSessionRepo) Delete(ctx context.Context, id uuid.UUID) error {
	session, ok := r.sessionsByID[id]
	if ok {
		delete(r.sessionsByToken, session.RefreshToken)
		delete(r.sessionsByID, id)
	}
	return nil
}

func (r *fakeSessionRepo) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	for id, session := range r.sessionsByID {
		if session.UserID == userID {
			delete(r.sessionsByToken, session.RefreshToken)
			delete(r.sessionsByID, id)
		}
	}
	return nil
}

func TestAuthService_RegisterAndLogin(t *testing.T) {
	userRepo := newFakeUserRepo()
	sessionRepo := newFakeSessionRepo()
	jwtService := security.NewJWTService("test-secret", 1, 24)
	passService := security.NewPasswordService()

	svc := service.NewAuthService(userRepo, sessionRepo, jwtService, passService)

	user, err := svc.RegisterUser(context.Background(), "user@example.com", "password123", "User")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user.Email != "user@example.com" {
		t.Fatalf("expected email to match")
	}

	_, tokens, err := svc.AuthenticateUser(context.Background(), "user@example.com", "password123", "agent", "127.0.0.1")
	if err != nil {
		t.Fatalf("expected login success, got %v", err)
	}
	if tokens.AccessToken == "" || tokens.RefreshToken == "" {
		t.Fatalf("expected tokens to be generated")
	}
}

func TestAuthService_LoginInvalidPassword(t *testing.T) {
	userRepo := newFakeUserRepo()
	sessionRepo := newFakeSessionRepo()
	jwtService := security.NewJWTService("test-secret", 1, 24)
	passService := security.NewPasswordService()

	svc := service.NewAuthService(userRepo, sessionRepo, jwtService, passService)

	_, err := svc.RegisterUser(context.Background(), "user@example.com", "password123", "User")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, _, err = svc.AuthenticateUser(context.Background(), "user@example.com", "wrong", "agent", "127.0.0.1")
	if err == nil {
		t.Fatalf("expected error for invalid password")
	}
}

func TestAuthService_RefreshTokenRotation(t *testing.T) {
	userRepo := newFakeUserRepo()
	sessionRepo := newFakeSessionRepo()
	jwtService := security.NewJWTService("test-secret", 1, 24)
	passService := security.NewPasswordService()

	svc := service.NewAuthService(userRepo, sessionRepo, jwtService, passService)

	user, err := svc.RegisterUser(context.Background(), "user@example.com", "password123", "User")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, tokens, err := svc.AuthenticateUser(context.Background(), "user@example.com", "password123", "agent", "127.0.0.1")
	if err != nil {
		t.Fatalf("expected login success, got %v", err)
	}

	oldRefresh := tokens.RefreshToken
	newTokens, err := svc.RefreshAccessToken(context.Background(), oldRefresh)
	if err != nil {
		t.Fatalf("expected refresh success, got %v", err)
	}
	if newTokens.RefreshToken == oldRefresh {
		t.Fatalf("expected refresh token rotation")
	}

	oldSession, err := sessionRepo.FindByRefreshToken(context.Background(), oldRefresh)
	if err != nil {
		t.Fatalf("expected old session to exist")
	}
	if oldSession.RevokedAt == nil {
		t.Fatalf("expected old session to be revoked")
	}

	newSession, err := sessionRepo.FindByRefreshToken(context.Background(), newTokens.RefreshToken)
	if err != nil {
		t.Fatalf("expected new session to exist")
	}
	if newSession.UserID != user.ID {
		t.Fatalf("expected session user to match")
	}
	if time.Now().After(newSession.ExpiresAt) {
		t.Fatalf("expected new session to be active")
	}
}

func TestAuthService_LogoutRevokesSession(t *testing.T) {
	userRepo := newFakeUserRepo()
	sessionRepo := newFakeSessionRepo()
	jwtService := security.NewJWTService("test-secret", 1, 24)
	passService := security.NewPasswordService()

	svc := service.NewAuthService(userRepo, sessionRepo, jwtService, passService)

	_, err := svc.RegisterUser(context.Background(), "user@example.com", "password123", "User")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, tokens, err := svc.AuthenticateUser(context.Background(), "user@example.com", "password123", "agent", "127.0.0.1")
	if err != nil {
		t.Fatalf("expected login success, got %v", err)
	}

	if err := svc.Logout(context.Background(), tokens.RefreshToken); err != nil {
		t.Fatalf("expected logout success, got %v", err)
	}

	session, err := sessionRepo.FindByRefreshToken(context.Background(), tokens.RefreshToken)
	if err != nil {
		t.Fatalf("expected session to exist")
	}
	if session.RevokedAt == nil {
		t.Fatalf("expected session to be revoked")
	}
}
