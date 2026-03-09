// Package dto defines Data Transfer Objects for users module.
//
// DTOs decouple API contracts from domain entities.

package dto

import (
	"mytodo/apps/api/internal/users/domain/entity"
	"time"

	"github.com/google/uuid"
)

// UserProfileDTO represents user profile data for API responses.
type UserProfileDTO struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Username    *string   `json:"username"`
	DisplayName string    `json:"display_name"`
	AvatarURL   string    `json:"avatar_url,omitempty"`
	Bio         string    `json:"bio,omitempty"`
	Location    string    `json:"location,omitempty"`
	Website     string    `json:"website,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Timezone    string    `json:"timezone"`
	Language    string    `json:"language"`
	Theme       string    `json:"theme"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// UserPreferencesDTO represents user preferences for API responses.
type UserPreferencesDTO struct {
	ID                     uuid.UUID `json:"id"`
	UserID                 uuid.UUID `json:"user_id"`
	EmailNotifications     bool      `json:"email_notifications"`
	PushNotifications      bool      `json:"push_notifications"`
	NewsletterSubscription bool      `json:"newsletter_subscription"`
	WeeklyDigest           bool      `json:"weekly_digest"`
	MentionsNotifications  bool      `json:"mentions_notifications"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

// ToUserProfileDTO converts a domain User entity to a UserProfileDTO.
func ToUserProfileDTO(user *entity.User) *UserProfileDTO {
	if user == nil {
		return nil
	}
	return &UserProfileDTO{
		ID:          user.ID,
		UserID:      user.UserID,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		AvatarURL:   user.AvatarURL,
		Bio:         user.Bio,
		Location:    user.Location,
		Website:     user.Website,
		Phone:       user.Phone,
		Timezone:    user.Timezone,
		Language:    user.Language,
		Theme:       user.Theme,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

// ToUserPreferencesDTO converts a domain Preference entity to UserPreferencesDTO.
func ToUserPreferencesDTO(pref *entity.Preference) *UserPreferencesDTO {
	if pref == nil {
		return nil
	}
	return &UserPreferencesDTO{
		ID:                     pref.ID,
		UserID:                 pref.UserID,
		EmailNotifications:     pref.EmailNotifications,
		PushNotifications:      pref.PushNotifications,
		NewsletterSubscription: pref.NewsletterSubscription,
		WeeklyDigest:           pref.WeeklyDigest,
		MentionsNotifications:  pref.MentionsNotifications,
		CreatedAt:              pref.CreatedAt,
		UpdatedAt:              pref.UpdatedAt,
	}
}

// PaginatedUserProfiles represents a paginated list of user profiles.
type PaginatedUserProfiles struct {
	Users []*UserProfileDTO `json:"users"`
	Total int               `json:"total"`
	Page  int               `json:"page"`
	Limit int               `json:"limit"`
}
