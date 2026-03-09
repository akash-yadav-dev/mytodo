// Package commands implements Command pattern for write operations.
package commands

import "errors"

// DeleteUserCommand represents a request to delete/deactivate a user.
type DeleteUserCommand struct {
	UserID     string `json:"user_id" validate:"required"`
	SoftDelete bool   `json:"soft_delete"` // true for deactivation
}

func (c DeleteUserCommand) Validate() error {
	if c.UserID == "" {
		return errors.New("user_id is required")
	}
	return nil
}
