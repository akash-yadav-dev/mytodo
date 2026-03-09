// Package dto defines Data Transfer Objects for users module.

package dto

// CreateUserProfileRequest is the request DTO for creating a user profile
type CreateUserProfileRequest struct {
	Username    string `json:"username" validate:"omitempty,min=3,max=30"`
	DisplayName string `json:"display_name" validate:"required,min=1,max=100"`
	AvatarURL   string `json:"avatar_url,omitempty" validate:"omitempty,url"`
}

// UpdateUserProfileRequest is the request DTO for updating a user profile
type UpdateUserProfileRequest struct {
	Username    *string `json:"username,omitempty" validate:"omitempty,min=3,max=30"`
	DisplayName string  `json:"display_name,omitempty" validate:"omitempty,min=1,max=100"`
	Bio         *string `json:"bio,omitempty" validate:"omitempty,max=500"`
	Location    *string `json:"location,omitempty" validate:"omitempty,max=100"`
	Website     *string `json:"website,omitempty" validate:"omitempty,url"`
	AvatarURL   *string `json:"avatar_url,omitempty" validate:"omitempty,url"`
	Phone       *string `json:"phone,omitempty" validate:"omitempty,max=20"`
	Timezone    *string `json:"timezone,omitempty"`
	Language    *string `json:"language,omitempty" validate:"omitempty,len=2"`
	Theme       *string `json:"theme,omitempty" validate:"omitempty,oneof=light dark"`
}

// UpdateUserPreferencesRequest is the request DTO for updating user preferences
type UpdateUserPreferencesRequest struct {
	EmailNotifications     bool `json:"email_notifications"`
	PushNotifications      bool `json:"push_notifications"`
	NewsletterSubscription bool `json:"newsletter_subscription"`
	WeeklyDigest           bool `json:"weekly_digest"`
	MentionsNotifications  bool `json:"mentions_notifications"`
}
