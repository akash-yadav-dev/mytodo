// Package queries implements Query pattern for project read operations.

package queries

// ListProjectsQuery retrieves projects with filters.
//
// Example structure:
//   type ListProjectsQuery struct {
//       UserID     string            `json:"user_id"` // Projects user has access to
//       Page       int               `json:"page"`
//       Limit      int               `json:"limit"`
//       Status     string            `json:"status"` // active, archived
//       Visibility string            `json:"visibility"`
//       Filters    map[string]string `json:"filters"`
//   }
//
// Example usage:
//   query := ListProjectsQuery{UserID: "user-123", Status: "active", Page: 1, Limit: 20}
//   result, err := handler.Handle(query)
//   // Returns: &PaginatedProjects{Projects: []ProjectDTO{...}, Total: 45}, nil
