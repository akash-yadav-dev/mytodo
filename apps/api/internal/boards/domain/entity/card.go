// Package entity defines core domain entities for the boards module.

package entity

// Card represents a card (issue/task) displayed on a board.
//
// In production, Card entities include:
// - ID, IssueID (reference to actual issue)
// - ColumnID, Position
// - Display metadata
//
// Example structure:
//   type Card struct {
//       ID       string `json:"id"`
//       IssueID  string `json:"issue_id"`
//       ColumnID string `json:"column_id"`
//       Position int    `json:"position"`
//   }
