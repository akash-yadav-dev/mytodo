// Package repository defines data access interfaces for comments.

package repository

// CommentRepository defines data access methods for comments.
//
// Example interface:
//   type CommentRepository interface {
//       Create(comment *Comment) error
//       FindByID(id string) (*Comment, error)
//       FindByEntity(entityID string) ([]Comment, error)
//       Update(comment *Comment) error
//       Delete(id string) error
//   }
//
// Example usage:
//   comments, _ := repo.FindByEntity("issue-123")
//   // Returns: []Comment{{ID: "1", Content: "..."},...}
