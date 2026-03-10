// Package service contains domain services for the users module.
//
// This file implements profile management business logic.

package service

import (
	"context"
	"errors"
	"strings"

	"mytodo/apps/api/internal/users/domain/entity"
	"mytodo/apps/api/internal/users/domain/repository"

	"github.com/google/uuid"
)

// ProfileService handles user profile write operations.
type ProfileService interface {
	CreateProfile(ctx context.Context, profile *entity.User) error
	UpdateProfile(ctx context.Context, profile *entity.User) error
	DeleteProfile(ctx context.Context, userID uuid.UUID) error
	GetProfileByUserID(ctx context.Context, userID uuid.UUID) (*entity.User, error)
	CreatePreferences(ctx context.Context, pref *entity.Preference) error
	GetPreferencesByUserID(ctx context.Context, userID uuid.UUID) (*entity.Preference, error)
	UpdatePreferences(ctx context.Context, pref *entity.Preference) error
	IsUsernameAvailable(ctx context.Context, username string) (bool, error)
}

type ProfileServiceImpl struct {
	userRepo repository.UserRepository
}

func NewProfileService(userRepo repository.UserRepository) ProfileService {
	return &ProfileServiceImpl{userRepo: userRepo}
}

func (s *ProfileServiceImpl) CreateProfile(ctx context.Context, profile *entity.User) error {
	if profile == nil {
		return errors.New("profile is required")
	}
	return s.userRepo.CreateProfile(ctx, profile)
}

func (s *ProfileServiceImpl) UpdateProfile(ctx context.Context, profile *entity.User) error {
	if profile == nil {
		return errors.New("profile is required")
	}
	return s.userRepo.UpdateProfile(ctx, profile)
}

func (s *ProfileServiceImpl) DeleteProfile(ctx context.Context, userID uuid.UUID) error {
	if userID == uuid.Nil {
		return errors.New("user id is required")
	}
	return s.userRepo.DeleteProfile(ctx, userID)
}

func (s *ProfileServiceImpl) GetProfileByUserID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	if userID == uuid.Nil {
		return nil, errors.New("user id is required")
	}
	return s.userRepo.FindProfileByUserID(ctx, userID)
}

func (s *ProfileServiceImpl) CreatePreferences(ctx context.Context, pref *entity.Preference) error {
	if pref == nil {
		return errors.New("preferences are required")
	}
	return s.userRepo.CreatePreferences(ctx, pref)
}

func (s *ProfileServiceImpl) GetPreferencesByUserID(ctx context.Context, userID uuid.UUID) (*entity.Preference, error) {
	if userID == uuid.Nil {
		return nil, errors.New("user id is required")
	}
	return s.userRepo.FindPreferencesByUserID(ctx, userID)
}

func (s *ProfileServiceImpl) UpdatePreferences(ctx context.Context, pref *entity.Preference) error {
	if pref == nil {
		return errors.New("preferences are required")
	}
	return s.userRepo.UpdatePreferences(ctx, pref)
}

func (s *ProfileServiceImpl) IsUsernameAvailable(ctx context.Context, username string) (bool, error) {
	username = strings.TrimSpace(username)
	if username == "" {
		return false, errors.New("username is required")
	}
	exists, err := s.userRepo.ExistsByUsername(ctx, username)
	if err != nil {
		return false, err
	}
	return !exists, nil
}
