// Package commands implements Command pattern for write operations.
//
// Commands represent user modification operations in the CQRS pattern.

package commands

// CreateUserCommand represents a request to create a new user.
//
// Example structure:
//   type CreateUserCommand struct {
//       Email       string `json:"email" validate:"required,email"`
//       Username    string `json:"username" validate:"required,min=3,max=30"`
//       DisplayName string `json:"display_name" validate:"required"`
//       Avatar      string `json:"avatar,omitempty"`
//   }
//
// Example usage:
//   cmd := CreateUserCommand{
//       Email:       "john@example.com",
//       Username:    "johndoe",
//       DisplayName: "John Doe",
//   }
//   result, err := handler.Handle(cmd)
//   // Returns: &User{ID: "uuid", Email: "john@example.com",...}, nil
