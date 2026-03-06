// Package commands implements Command pattern for project write operations.

package commands

// ArchiveProjectCommand archives a project for later restoration.
//
// Example structure:
//   type ArchiveProjectCommand struct {
//       ProjectID string `json:"project_id" validate:"required"`
//   }
