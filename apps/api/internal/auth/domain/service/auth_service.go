// Package service contains domain services that implement business logic
// spanning multiple entities or requiring coordination between entities.
//
// Domain services in Clean Architecture represent business operations that:
// - Don't naturally belong to a single entity
// - Coordinate operations across multiple entities
// - Implement complex business rules and workflows
// - Remain independent of infrastructure concerns
//
// In production applications, this folder typically contains:
// - Business logic that requires multiple repositories
// - Domain calculations and validations
// - Transaction coordination (at domain level)
// - Domain event publishing logic
// - Business workflow orchestration

package service

import (
	"context"
	"errors"
	"mytodo/apps/api/internal/auth/domain/entity"
	"mytodo/apps/api/internal/auth/domain/repository"
	"mytodo/apps/api/pkg/security"
	"time"

	"github.com/google/uuid"
)

// AuthService handles core authentication business logic.
// In production applications, auth services typically implement:
// - User registration with validation
// - Login authentication (password verification, account checks)
// - Token generation and validation (JWT, refresh tokens)
// - Password reset workflows
// - Email/phone verification
// - Account activation/deactivation
// - Multi-factor authentication (MFA) setup and verification
// - OAuth integration and social login
// - Brute force protection and rate limiting
// - Audit logging of authentication events

type AuthService struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	jwtService  *security.JWTService
	passService *security.PasswordService
}

type AuthTokens struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}

func NewAuthService(
	userRepo repository.UserRepository,
	sessionRepo repository.SessionRepository,
	jwtService *security.JWTService,
	passService *security.PasswordService,
) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		jwtService:  jwtService,
		passService: passService,
	}
}

// RegisterUser creates a new user account
func (s *AuthService) RegisterUser(ctx context.Context, email, password, name string) (*entity.User, error) {
	// Check if user already exists
	exists, err := s.userRepo.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := s.passService.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := entity.NewUser(email, name, hashedPassword)

	// Save to database
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// AuthenticateUser verifies credentials and returns tokens
func (s *AuthService) AuthenticateUser(ctx context.Context, email, password, userAgent, ipAddress string) (*entity.User, *AuthTokens, error) {
	// Find user by email
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, nil, errors.New("invalid credentials")
	}

	// Check if account is active
	if !user.IsActive {
		return nil, nil, errors.New("account is deactivated")
	}

	// Verify password
	if !s.passService.VerifyPassword(password, user.PasswordHash) {
		return nil, nil, errors.New("invalid credentials")
	}

	// Generate tokens
	tokens, err := s.generateTokens(ctx, user, userAgent, ipAddress)
	if err != nil {
		return nil, nil, err
	}

	// Update last login
	user.UpdateLastLogin()
	if err := s.userRepo.Update(ctx, user); err != nil {
		// Log error but don't fail authentication
		// TODO: Add proper logging
	}

	return user, tokens, nil
}

// RefreshAccessToken generates a new access token using a refresh token
func (s *AuthService) RefreshAccessToken(ctx context.Context, refreshToken string) (*AuthTokens, error) {
	// Validate refresh token
	claims, err := s.jwtService.ValidateToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// Find session
	session, err := s.sessionRepo.FindByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, errors.New("session not found")
	}

	// Check if session is expired or revoked
	if session.IsExpired() {
		return nil, errors.New("session expired")
	}
	if session.IsRevoked() {
		return nil, errors.New("session revoked")
	}

	// Get user
	user, err := s.userRepo.FindByID(ctx, claims.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Revoke current session before issuing a new refresh token
	session.Revoke()
	if err := s.sessionRepo.Update(ctx, session); err != nil {
		return nil, err
	}

	// Generate new access token
	accessToken, err := s.jwtService.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	// Rotate refresh token
	newRefreshToken, err := s.jwtService.GenerateRefreshToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	// Store new session
	newExpiresAt := time.Now().Add(s.jwtService.GetRefreshTokenExpiration())
	newSession := entity.NewSession(user.ID, newRefreshToken, session.UserAgent, session.IPAddress, newExpiresAt)
	if err := s.sessionRepo.Create(ctx, newSession); err != nil {
		return nil, err
	}

	return &AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    3600, // 1 hour in seconds
	}, nil
}

// Logout revokes the user's session
func (s *AuthService) Logout(ctx context.Context, refreshToken string) error {
	// Find session
	session, err := s.sessionRepo.FindByRefreshToken(ctx, refreshToken)
	if err != nil {
		// Session not found is not an error for logout
		return nil
	}

	// Revoke session
	session.Revoke()
	return s.sessionRepo.Update(ctx, session)
}

// generateTokens creates access and refresh tokens and stores the session
func (s *AuthService) generateTokens(ctx context.Context, user *entity.User, userAgent, ipAddress string) (*AuthTokens, error) {
	// Generate access token
	accessToken, err := s.jwtService.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshToken, err := s.jwtService.GenerateRefreshToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	// Create session
	expiresAt := time.Now().Add(s.jwtService.GetRefreshTokenExpiration())
	session := entity.NewSession(user.ID, refreshToken, userAgent, ipAddress, expiresAt)

	// Save session
	if err := s.sessionRepo.Create(ctx, session); err != nil {
		return nil, err
	}

	return &AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    3600, // 1 hour in seconds
	}, nil
}

// GetUserByID retrieves user information
func (s *AuthService) GetUserByID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	return s.userRepo.FindByID(ctx, userID)
}
