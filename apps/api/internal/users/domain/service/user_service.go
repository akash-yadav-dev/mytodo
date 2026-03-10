// Package service contains domain services for the users module.
//
// Domain services implement business logic that doesn't naturally fit
// within a single entity and coordinates operations across entities.

package service

import (
	"context"
	"errors"
	"mytodo/apps/api/internal/users/domain/entity"
	"mytodo/apps/api/internal/users/domain/repository"

	"github.com/google/uuid"
)

type UserService interface {
	GetProfileByUserID(ctx context.Context, userID uuid.UUID) (*entity.User, error)
	GetProfileByID(ctx context.Context, profileID uuid.UUID) (*entity.User, error)
	ListProfiles(ctx context.Context, page, limit int) ([]*entity.User, int, error)
	SearchProfiles(ctx context.Context, query string, limit int) ([]*entity.User, error)
	GetPreferencesByUserID(ctx context.Context, userID uuid.UUID) (*entity.Preference, error)
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) GetProfileByUserID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	if userID == uuid.Nil {
		return nil, errors.New("user id is required")
	}
	return s.userRepo.FindProfileByUserID(ctx, userID)
}

func (s *UserServiceImpl) GetProfileByID(ctx context.Context, profileID uuid.UUID) (*entity.User, error) {
	if profileID == uuid.Nil {
		return nil, errors.New("profile id is required")
	}
	return s.userRepo.FindProfileByID(ctx, profileID)
}

func (s *UserServiceImpl) ListProfiles(ctx context.Context, page, limit int) ([]*entity.User, int, error) {
	return s.userRepo.ListProfiles(ctx, page, limit)
}

func (s *UserServiceImpl) SearchProfiles(ctx context.Context, query string, limit int) ([]*entity.User, error) {
	return s.userRepo.SearchProfiles(ctx, query, limit)
}

func (s *UserServiceImpl) GetPreferencesByUserID(ctx context.Context, userID uuid.UUID) (*entity.Preference, error) {
	if userID == uuid.Nil {
		return nil, errors.New("user id is required")
	}
	return s.userRepo.FindPreferencesByUserID(ctx, userID)
}
