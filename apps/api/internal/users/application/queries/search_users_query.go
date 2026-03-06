// Package queries implements Query pattern for read operations.

package queries

// SearchUsersQuery performs full-text search across user fields.
//
// Example structure:
//   type SearchUsersQuery struct {
//       Query  string   `json:"query" validate:"required,min=2"`
//       Fields []string `json:"fields"` // ["email", "username", "display_name"]
//       Limit  int      `json:"limit" validate:"max=50"`
//   }
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
