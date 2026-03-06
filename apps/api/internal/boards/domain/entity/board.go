// Package entity defines core domain entities for the boards module.
//
// Boards provide kanban-style visual organization of work items.

package entity

// Board represents a kanban board for visual work organization.
//
// In production, Board entities typically include:
// - ID, Name, ProjectID
// - Columns (ordered list)
// - Settings (WIP limits, card styles)
// - Visibility and permissions
//
// Example structure:
//   type Board struct {
//       ID        string   `json:"id"`
//       Name      string   `json:"name"`
//       ProjectID string   `json:"project_id"`
//       Columns   []Column `json:"columns"`
//       CreatedAt time.Time `json:"created_at"`
//   }
//
// Example methods:
//   func (b *Board) AddColumn(name string, position int) error
//   func (b *Board) MoveCard(cardID, fromColumn, toColumn string) error
