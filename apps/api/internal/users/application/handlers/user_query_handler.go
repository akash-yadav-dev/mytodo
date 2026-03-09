// Package handlers orchestrates use cases for the users module.

package handlers

import (
	"context"
	"errors"
	"fmt"
	"mytodo/apps/api/internal/users/application/queries"
	"mytodo/apps/api/internal/users/domain/repository"
	"mytodo/apps/api/internal/users/interfaces/dto"
	"strings"

	"github.com/google/uuid"
)

// UserQueryHandler processes user read operations.
type UserQueryHandler struct {
	userRepo repository.UserRepository
}

// NewUserQueryHandler creates a new UserQueryHandler
func NewUserQueryHandler(userRepo repository.UserRepository) *UserQueryHandler {
	return &UserQueryHandler{
		userRepo: userRepo,
	}
}

// HandleGetProfileByID retrieves a user profile by profile ID
func (h *UserQueryHandler) HandleGetProfileByID(ctx context.Context, query queries.GetUserByIDQuery) (*dto.UserProfileDTO, error) {
	// Step 1: Validate query
	if err := query.Validate(); err != nil {
		return nil, err
	}

	// Step 2: Parse UUID
	userID, err := uuid.Parse(query.UserID)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	// Step 3: Fetch user profile from repository
	user, err := h.userRepo.FindProfileByUserID(ctx, userID)
	if err != nil {
		return nil, errors.New("user profile not found")
	}

	// Step 4: Convert entity to DTO
	return dto.ToUserProfileDTO(user), nil
}

// HandleGetUserProfileList retrieves a paginated list of user profiles
func (h *UserQueryHandler) HandleGetUserProfileList(ctx context.Context, query queries.ListUsersQuery) (*dto.PaginatedUserProfiles, error) {
	// Step 1: Validate query
	var validationErrors []string

	if query.Page < 1 {
		validationErrors = append(validationErrors, "page must be greater than 0")
	}

	if query.Limit < 1 || query.Limit > 100 {
		validationErrors = append(validationErrors, "limit must be between 1 and 100")
	}

	if len(validationErrors) > 0 {
		return nil, fmt.Errorf("validation failed: %s", strings.Join(validationErrors, ", "))
	}

	// Step 2: Fetch users from repository
	users, total, err := h.userRepo.ListProfiles(ctx, query.Page, query.Limit)
	if err != nil {
		return nil, err
	}

	// Step 3: Convert entities to DTOs
	userDTOs := make([]*dto.UserProfileDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dto.ToUserProfileDTO(user)
	}

	// Step 4: Return paginated result
	return &dto.PaginatedUserProfiles{
		Users: userDTOs,
		Total: total,
		Page:  query.Page,
		Limit: query.Limit,
	}, nil
}

// HandleSearchUserProfiles searches for user profiles
func (h *UserQueryHandler) HandleSearchUserProfiles(ctx context.Context, query queries.SearchUsersQuery) ([]*dto.UserProfileDTO, error) {
	// Step 1: Validate query
	if err := query.Validate(); err != nil {
		return nil, err
	}

	// Step 2: Search users
	users, err := h.userRepo.SearchProfiles(ctx, query.Query, query.Limit)
	if err != nil {
		return nil, err
	}

	// Step 3: Convert to DTOs
	userDTOs := make([]*dto.UserProfileDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dto.ToUserProfileDTO(user)
	}

	return userDTOs, nil
}

// HandleGetUserPreferences retrieves user preferences
func (h *UserQueryHandler) HandleGetUserPreferences(ctx context.Context, userID uuid.UUID) (*dto.UserPreferencesDTO, error) {
	prefs, err := h.userRepo.FindPreferencesByUserID(ctx, userID)
	if err != nil {
		return nil, errors.New("user preferences not found")
	}

	return dto.ToUserPreferencesDTO(prefs), nil
}
