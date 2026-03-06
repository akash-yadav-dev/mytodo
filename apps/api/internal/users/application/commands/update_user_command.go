// Package commands implements Command pattern for write operations.

package commands

// UpdateUserCommand represents a request to update user information.
//
// Example structure:
//   type UpdateUserCommand struct {
//       UserID      string `json:"user_id" validate:"required"`
//       DisplayName string `json:"display_name,omitempty"`
//       Avatar      string `json:"avatar,omitempty"`
//       Bio         string `json:"bio,omitempty"`
//   }
//
// Example usage:
//   cmd := UpdateUserCommand{
//       UserID:      "user-123",
//       DisplayName: "Jane Doe",
//       Bio:         "Software Engineer",
//   }
//   result, err := handler.Handle(cmd)
//   // Returns: &User{ID: "user-123", DisplayName: "Jane Doe",...}, nil
