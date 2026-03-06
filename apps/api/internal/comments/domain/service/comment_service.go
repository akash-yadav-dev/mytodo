// Package service contains domain services for the comments module.

package service

// CommentService handles comment business logic.
//
// Example interface:
//   type CommentService interface {
//       CreateComment(entityID, entityType, authorID, content string) (*Comment, error)
//       UpdateComment(commentID, newContent string) (*Comment, error)
//       DeleteComment(commentID string) error
//       GetComments(entityID string, includeThreads bool) ([]Comment, error)
//   }
//
// Example usage:
//   comment, _ := svc.CreateComment("issue-123", "issue", "user-1", "Looking good!")
//   // Returns: &Comment{ID: "comment-1", Content: "Looking good!",...}
