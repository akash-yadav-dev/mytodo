// Package queries implements Query pattern for project read operations.

package queries

// GetProjectQuery retrieves a single project.
//
// Example structure:
//   type GetProjectQuery struct {
//       ProjectID       string `json:"project_id" validate:"required"`
//       IncludeMembers  bool   `json:"include_members"`
//       IncludeConfig   bool   `json:"include_config"`
//   }
//
// Example usage:
//   query := GetProjectQuery{ProjectID: "proj-1", IncludeMembers: true}
//   result, err := handler.Handle(query)
//   // Returns: &ProjectDTO{ID: "proj-1", Members: [...],...}, nil
