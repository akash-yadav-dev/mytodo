// Package entity defines core domain entities for the boards module.

package entity

// Column represents a column in a kanban board.
//
// In production, Column entities include:
// - ID, Name, BoardID
// - Position (order)
// - WIPLimit (work-in-progress limit)
// - Cards (issues/tasks in this column)
//
// Example structure:
//   type Column struct {
//       ID       string `json:"id"`
//       Name     string `json:"name"`
//       BoardID  string `json:"board_id"`
//       Position int    `json:"position"`
//       WIPLimit int    `json:"wip_limit"`
//   }
