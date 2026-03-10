// Package handlers orchestrates use cases for the users module.
//
// Handlers coordinate domain services and repositories to execute commands.

package handlers

import (
	"context"
	"errors"
	"mytodo/apps/api/internal/users/application/commands"
	"mytodo/apps/api/internal/users/domain/entity"
	"mytodo/apps/api/internal/users/domain/service"
	"mytodo/apps/api/internal/users/interfaces/dto"

	"github.com/google/uuid"
)

// UserCommandHandler processes user write operations.
type UserCommandHandler struct {
	profileService service.ProfileService
}

// NewUserCommandHandler creates a new UserCommandHandler
func NewUserCommandHandler(profileService service.ProfileService) *UserCommandHandler {
	return &UserCommandHandler{
		profileService: profileService,
	}
}

// HandleCreateUserProfile creates a new user profile (called after user registration)
func (h *UserCommandHandler) HandleCreateUserProfile(ctx context.Context, cmd commands.CreateUserProfileCommand) (*dto.UserProfileDTO, error) {
	// Step 1: Validate command
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	// Step 2: Parse user ID
	authUserID, err := uuid.Parse(cmd.AuthUserID)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	// Step 3: Check if username is already taken (if provided)
	if cmd.Username != "" {
		available, err := h.profileService.IsUsernameAvailable(ctx, cmd.Username)
		if err != nil {
			return nil, err
		}
		if !available {
			return nil, errors.New("username already taken")
		}
	}

	// Step 4: Create user profile entity
	userProfile := entity.NewUserProfile(authUserID, cmd.DisplayName)
	if cmd.Username != "" {
		if err := userProfile.SetUsername(cmd.Username); err != nil {
			return nil, err
		}
	}
	if cmd.AvatarURL != "" {
		userProfile.AvatarURL = cmd.AvatarURL
	}

	// Step 5: Save to repository
	if err := h.profileService.CreateProfile(ctx, userProfile); err != nil {
		return nil, err
	}

	// Step 6: Create default preferences
	preferences := entity.NewUserPreferences(authUserID)
	if err := h.profileService.CreatePreferences(ctx, preferences); err != nil {
		// Log error but don't fail profile creation
		// TODO: Add proper logging
	}

	// Step 7: Return DTO
	return dto.ToUserProfileDTO(userProfile), nil
}

// HandleUpdateUserProfile updates user profile information
func (h *UserCommandHandler) HandleUpdateUserProfile(ctx context.Context, cmd commands.UpdateUserProfileCommand) (*dto.UserProfileDTO, error) {
	// Step 1: Validate command
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	// Step 2: Parse user ID
	userID, err := uuid.Parse(cmd.UserID)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	// Step 3: Fetch existing profile
	userProfile, err := h.profileService.GetProfileByUserID(ctx, userID)
	if err != nil {
		return nil, errors.New("user profile not found")
	}

	// Step 4: Check if new username is already taken
	if cmd.Username != nil && *cmd.Username != "" {
		if userProfile.Username == nil || *userProfile.Username != *cmd.Username {
			available, err := h.profileService.IsUsernameAvailable(ctx, *cmd.Username)
			if err != nil {
				return nil, err
			}
			if !available {
				return nil, errors.New("username already taken")
			}
			if err := userProfile.SetUsername(*cmd.Username); err != nil {
				return nil, err
			}
		}
	}

	// Step 5: Update profile fields
	if cmd.DisplayName != "" {
		userProfile.DisplayName = cmd.DisplayName
	}
	if cmd.Bio != nil {
		userProfile.Bio = *cmd.Bio
	}
	if cmd.Location != nil {
		userProfile.Location = *cmd.Location
	}
	if cmd.Website != nil {
		userProfile.Website = *cmd.Website
	}
	if cmd.AvatarURL != nil {
		userProfile.AvatarURL = *cmd.AvatarURL
	}
	if cmd.Phone != nil {
		userProfile.Phone = *cmd.Phone
	}

	// Step 6: Update preferences if provided
	if cmd.Timezone != nil || cmd.Language != nil || cmd.Theme != nil {
		timezone := userProfile.Timezone
		language := userProfile.Language
		theme := userProfile.Theme

		if cmd.Timezone != nil {
			timezone = *cmd.Timezone
		}
		if cmd.Language != nil {
			language = *cmd.Language
		}
		if cmd.Theme != nil {
			theme = *cmd.Theme
		}

		userProfile.UpdatePreferences(timezone, language, theme)
	}

	// Step 7: Save to repository
	if err := h.profileService.UpdateProfile(ctx, userProfile); err != nil {
		return nil, err
	}

	// Step 8: Return updated DTO
	return dto.ToUserProfileDTO(userProfile), nil
}

// HandleDeleteUserProfile deletes a user profile
func (h *UserCommandHandler) HandleDeleteUserProfile(ctx context.Context, cmd commands.DeleteUserCommand) error {
	// Step 1: Validate command
	if err := cmd.Validate(); err != nil {
		return err
	}

	// Step 2: Parse user ID
	userID, err := uuid.Parse(cmd.UserID)
	if err != nil {
		return errors.New("invalid user ID format")
	}

	// Step 3: Check if profile exists
	_, err = h.profileService.GetProfileByUserID(ctx, userID)
	if err != nil {
		return errors.New("user profile not found")
	}

	// Step 4: Delete profile (preferences will be cascade deleted)
	if err := h.profileService.DeleteProfile(ctx, userID); err != nil {
		return err
	}

	return nil
}

// HandleUpdateUserPreferences updates user notification preferences
func (h *UserCommandHandler) HandleUpdateUserPreferences(ctx context.Context, cmd commands.UpdateUserPreferencesCommand) (*dto.UserPreferencesDTO, error) {
	// Step 1: Validate command
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	// Step 2: Parse user ID
	userID, err := uuid.Parse(cmd.UserID)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	// Step 3: Fetch existing preferences
	prefs, err := h.profileService.GetPreferencesByUserID(ctx, userID)
	if err != nil {
		return nil, errors.New("user preferences not found")
	}

	// Step 4: Update preferences
	prefs.UpdateNotificationSettings(
		cmd.EmailNotifications,
		cmd.PushNotifications,
		cmd.NewsletterSubscription,
		cmd.WeeklyDigest,
		cmd.MentionsNotifications,
	)

	// Step 5: Save to repository
	if err := h.profileService.UpdatePreferences(ctx, prefs); err != nil {
		return nil, err
	}

	// Step 6: Return updated DTO
	return dto.ToUserPreferencesDTO(prefs), nil
}
