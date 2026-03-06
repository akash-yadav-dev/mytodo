// Package queries implements Query pattern for read operations.

package queries

// GetUserByEmailQuery retrieves a user by their email address.
//
// Example structure:
//   type GetUserByEmailQuery struct {
//       Email string `json:"email" validate:"required,email"`
//   }
//
// Example usage:
//   query := GetUserByEmailQuery{Email: "john@example.com"}
//   result, err := handler.Handle(query)
//   // Returns: &UserDTO{ID: "user-123", Email: "john@example.com",...}, nil
//   // Returns: nil, ErrUserNotFound if user doesn't exist
