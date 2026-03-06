// Package entity defines core domain entities for the comments module.
//
// Comments allow discussion and collaboration on issues and tasks.

package entity

// Comment represents a comment on an issue or task.
//
// In production, Comment entities include:
// - ID, Content, AuthorID
// - ParentID (for threaded comments)
// - EntityID, EntityType (issue, PR, etc.)
// - Timestamps, EditHistory
// - Reactions and mentions
//
// Example structure:
//   type Comment struct {
//       ID         string    `json:"id"`
//       Content    string    `json:"content"`
//       AuthorID   string    `json:"author_id"`
//       EntityID   string    `json:"entity_id"`
//       EntityType string    `json:"entity_type"`
//       ParentID   string    `json:"parent_id,omitempty"`
//       CreatedAt  time.Time `json:"created_at"`
//       UpdatedAt  time.Time `json:"updated_at"`
//   }
//
// Example methods:
//   func (c *Comment) Edit(newContent string) error
//   func (c *Comment) IsEdited() bool
