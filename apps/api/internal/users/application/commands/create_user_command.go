// Package commands implements Command pattern for write operations.
//
// Commands represent user modification operations in the CQRS pattern.

package commands

import "errors"

// CreateUserProfileCommand represents a request to create a new user profile.
type CreateUserProfileCommand struct {
	AuthUserID  string `json:"auth_user_id" validate:"required"`
	Username    string `json:"username" validate:"omitempty,min=3,max=30"`
	DisplayName string `json:"display_name" validate:"required,min=1,max=100"`
	AvatarURL   string `json:"avatar_url,omitempty"`
}

// Validate validates the create user profile command
func (c CreateUserProfileCommand) Validate() error {
	if c.AuthUserID == "" {
		return errors.New("auth_user_id is required")
	}
	if c.DisplayName == "" {
		return errors.New("display_name is required")
	}
	if len(c.DisplayName) > 100 {
		return errors.New("display_name must not exceed 100 characters")
	}
	if c.Username != "" {
		if len(c.Username) < 3 || len(c.Username) > 30 {
			return errors.New("username must be between 3 and 30 characters")
		}
	}
	return nil
}
