// Package queries implements Query pattern for read operations.

package queries

import "errors"

// SearchUsersQuery performs full-text search across user fields.
type SearchUsersQuery struct {
	Query  string   `json:"query" validate:"required,min=2"`
	Fields []string `json:"fields"` // ["email", "username", "display_name"]
	Limit  int      `json:"limit" validate:"max=50"`
}

// Validate validates the search query
func (q *SearchUsersQuery) Validate() error {
	if q.Query == "" {
		return errors.New("search query is required")
	}
	if len(q.Query) < 2 {
		return errors.New("search query must be at least 2 characters")
	}
	if q.Limit < 1 {
		q.Limit = 20
	}
	if q.Limit > 100 {
		return errors.New("limit cannot exceed 100")
	}
	return nil
}

//
// Example usage:
//   query := SearchUsersQuery{
//       Query:  "john",
//       Fields: []string{"username", "display_name"},
//       Limit:  10,
//   }
//   results, err := handler.Handle(query)
//   // Returns: []UserDTO{
//   //   {ID: "1", Username: "johndoe", DisplayName: "John Doe"},
//   //   {ID: "2", Username: "john_smith", DisplayName: "John Smith"},
//   // }, nil
