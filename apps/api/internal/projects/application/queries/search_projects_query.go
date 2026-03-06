// Package queries implements Query pattern for project read operations.

package queries

// SearchProjectsQuery performs full-text search on projects.
//
// Example structure:
//   type SearchProjectsQuery struct {
//       Query  string `json:"query" validate:"required,min=2"`
//       UserID string `json:"user_id"` // Limit to user's projects
//       Limit  int    `json:"limit"`
//   }
//
// Example usage:
//   query := SearchProjectsQuery{Query: "todo", UserID: "user-123", Limit: 10}
//   results, _ := handler.Handle(query)
//   // Returns: []ProjectDTO{{ID: "proj-1", Name: "My Todo",...}}
