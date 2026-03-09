// Package commands implements Command pattern for write operations.

package commands

import "errors"

// UpdateUserProfileCommand represents a request to update user profile information.
type UpdateUserProfileCommand struct {
	UserID      string  `json:"user_id" validate:"required"`
	Username    *string `json:"username,omitempty"`
	DisplayName string  `json:"display_name,omitempty"`
	Bio         *string `json:"bio,omitempty"`
	Location    *string `json:"location,omitempty"`
	Website     *string `json:"website,omitempty"`
	AvatarURL   *string `json:"avatar_url,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	Timezone    *string `json:"timezone,omitempty"`
	Language    *string `json:"language,omitempty"`
	Theme       *string `json:"theme,omitempty"`
}

// Validate validates the update user profile command
func (c UpdateUserProfileCommand) Validate() error {
	if c.UserID == "" {
		return errors.New("user_id is required")
	}
	if c.Username != nil && *c.Username != "" {
		if len(*c.Username) < 3 || len(*c.Username) > 30 {
			return errors.New("username must be between 3 and 30 characters")
		}
	}
	if c.DisplayName != "" && len(c.DisplayName) > 100 {
		return errors.New("display_name must not exceed 100 characters")
	}
	return nil
}

// UpdateUserPreferencesCommand represents a request to update user preferences.
type UpdateUserPreferencesCommand struct {
	UserID                 string `json:"user_id" validate:"required"`
	EmailNotifications     bool   `json:"email_notifications"`
	PushNotifications      bool   `json:"push_notifications"`
	NewsletterSubscription bool   `json:"newsletter_subscription"`
	WeeklyDigest           bool   `json:"weekly_digest"`
	MentionsNotifications  bool   `json:"mentions_notifications"`
}

// Validate validates the update user preferences command
func (c UpdateUserPreferencesCommand) Validate() error {
	if c.UserID == "" {
		return errors.New("user_id is required")
	}
	return nil
}
