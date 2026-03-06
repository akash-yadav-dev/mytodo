// Package commands implements Command pattern for boards write operations.

package commands

// CreateBoardCommand creates a new kanban board.
//
// Example structure:
//   type CreateBoardCommand struct {
//       ProjectID string   `json:"project_id" validate:"required"`
//       Name      string   `json:"name" validate:"required"`
//       Columns   []string `json:"columns"` // Initial column names
//   }
