// Package commands implements Command pattern for project write operations.

package commands

// DeleteProjectCommand deletes/archives a project.
//
// Example structure:
//   type DeleteProjectCommand struct {
//       ProjectID string `json:"project_id" validate:"required"`
//       HardDelete bool   `json:"hard_delete"` // true to permanently delete
//   }
