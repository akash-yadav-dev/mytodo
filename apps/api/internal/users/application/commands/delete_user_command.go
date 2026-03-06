// Package commands implements Command pattern for write operations.

package commands

// DeleteUserCommand represents a request to delete/deactivate a user.
//
// Example structure:
//   type DeleteUserCommand struct {
//       UserID   string `json:"user_id" validate:"required"`
//       SoftDelete bool   `json:"soft_delete"` // true for deactivation
//   }
//
// Example usage:
//   cmd := DeleteUserCommand{UserID: "user-123", SoftDelete: true}
//   err := handler.Handle(cmd)
//   // Returns: nil on success, error on failure
