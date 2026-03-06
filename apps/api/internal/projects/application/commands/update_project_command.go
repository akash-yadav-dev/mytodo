// Package commands implements Command pattern for project write operations.

package commands

// UpdateProjectCommand updates project information.
//
// Example structure:
//   type UpdateProjectCommand struct {
//       ProjectID   string `json:"project_id" validate:"required"`
//       Name        string `json:"name,omitempty"`
//       Description string `json:"description,omitempty"`
//       Visibility  string `json:"visibility,omitempty"`
//   }
