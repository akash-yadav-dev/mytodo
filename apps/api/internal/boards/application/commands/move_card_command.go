// Package commands implements Command pattern for boards write operations.

package commands

// MoveCardCommand moves a card between columns.
//
// Example structure:
//   type MoveCardCommand struct {
//       CardID     string `json:"card_id" validate:"required"`
//       ToColumnID string `json:"to_column_id" validate:"required"`
//       Position   int    `json:"position"`
//   }
