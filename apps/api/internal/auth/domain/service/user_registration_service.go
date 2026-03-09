// Package service contains domain services for auth module integration.

package service

import (
	"context"
	"mytodo/apps/api/internal/auth/domain/entity"
	authRepo "mytodo/apps/api/internal/auth/domain/repository"
	"mytodo/apps/api/internal/users/application/commands"
	"mytodo/apps/api/internal/users/application/handlers"
	"mytodo/apps/api/pkg/security"

	"github.com/google/uuid"
)

// UserRegistrationService handles user registration with profile creation
// This service coordinates between auth and users modules
type UserRegistrationService struct {
	authService        *AuthService
	userProfileHandler *handlers.UserCommandHandler
	userRepo           authRepo.UserRepository
	sessionRepo        authRepo.SessionRepository
	jwtService         *security.JWTService
	passwordService    *security.PasswordService
}

// NewUserRegistrationService creates a new UserRegistrationService
func NewUserRegistrationService(
	authService *AuthService,
	userProfileHandler *handlers.UserCommandHandler,
	userRepo authRepo.UserRepository,
	sessionRepo authRepo.SessionRepository,
	jwtService *security.JWTService,
	passwordService *security.PasswordService,
) *UserRegistrationService {
	return &UserRegistrationService{
		authService:        authService,
		userProfileHandler: userProfileHandler,
		userRepo:           userRepo,
		sessionRepo:        sessionRepo,
		jwtService:         jwtService,
		passwordService:    passwordService,
	}
}

// RegisterUserWithProfile registers a new user and creates their profile
func (s *UserRegistrationService) RegisterUserWithProfile(
	ctx context.Context,
	email, password, name, userAgent, ipAddress string,
) (*entity.User, *AuthTokens, error) {
	// Step 1: Register user in auth module
	user, err := s.authService.RegisterUser(ctx, email, password, name)
	if err != nil {
		return nil, nil, err
	}

	// Step 2: Create user profile in users module
	profileCmd := commands.CreateUserProfileCommand{
		AuthUserID:  user.ID.String(),
		DisplayName: name,
	}

	_, err = s.userProfileHandler.HandleCreateUserProfile(ctx, profileCmd)
	if err != nil {
		// Log error but don't fail registration
		// The user can create their profile later
		// TODO: Add proper logging
	}

	// Step 3: Authenticate user to get tokens
	_, tokens, err := s.authService.AuthenticateUser(
		ctx,
		email,
		password,
		userAgent,
		ipAddress,
	)
	if err != nil {
		return user, nil, err
	}

	return user, tokens, nil
}

// GetOrCreateUserProfile gets existing profile or creates a new one
func (s *UserRegistrationService) GetOrCreateUserProfile(ctx context.Context, userID uuid.UUID, displayName string) error {
	// Try to create profile (will fail if already exists)
	profileCmd := commands.CreateUserProfileCommand{
		AuthUserID:  userID.String(),
		DisplayName: displayName,
	}

	_, err := s.userProfileHandler.HandleCreateUserProfile(ctx, profileCmd)
	// Ignore errors - profile might already exist
	return err
}
