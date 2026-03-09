// Package queries implements Query pattern for read operations.
//
// Queries represent data retrieval operations in the CQRS pattern.
// Each query defines the data needed to perform a read operation.

package queries

import (
	"errors"
)

// GetUserByIDQuery retrieves a single user by their unique identifier.
//
// This query is typically handled by UserQueryHandler.HandleGetByID.
//
// Fields:
//   - UserID: Unique identifier of the user
//   - IncludeProfile: Indicates whether the user profile should be loaded
//
// Example usage:
//
//	query := GetUserByIDQuery{
//		UserID: "user-123",
//		IncludeProfile: true,
//	}
//
//	result, err := handler.HandleGetByID(query)
//
// Expected result:
//
//	&UserDTO{
//	    ID: "user-123",
//	    Email: "user@example.com",
//	    Name: "John Doe",
//	    Profile: {...},
//	}
type GetUserByIDQuery struct {
	UserID         string `json:"user_id" validate:"required"`
	IncludeProfile bool   `json:"include_profile"`
}

// Validate performs basic validation on the query before execution.
//
// This ensures invalid queries are rejected early in the application layer.
func (q GetUserByIDQuery) Validate() error {

	if q.UserID == "" {
		return errors.New("user_id is required")
	}

	// Future validations can be added here:
	// - UUID format check
	// - Length validation
	// - Security rules

	return nil
}

// NewGetUserByIDQuery creates a new query instance with validation.
//
// This is an optional helper constructor used in service/controller layers.
func NewGetUserByIDQuery(userID string, includeProfile bool) (*GetUserByIDQuery, error) {

	query := &GetUserByIDQuery{
		UserID:         userID,
		IncludeProfile: includeProfile,
	}

	if err := query.Validate(); err != nil {
		return nil, err
	}

	return query, nil
}
