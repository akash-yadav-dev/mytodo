// Package queries implements Query pattern for read operations.
//
// Queries represent data retrieval operations in the CQRS pattern.

package queries

// GetUserByIDQuery retrieves a single user by their unique identifier.
//
// Example structure:
//   type GetUserByIDQuery struct {
//       UserID         string `json:"user_id" validate:"required"`
//       IncludeProfile bool   `json:"include_profile"`
//   }
//
// Example usage:
//   query := GetUserByIDQuery{UserID: "user-123", IncludeProfile: true}
//   result, err := handler.Handle(query)
//   // Returns: &UserDTO{ID: "user-123", Email: "...", Profile: {...}}, nil
