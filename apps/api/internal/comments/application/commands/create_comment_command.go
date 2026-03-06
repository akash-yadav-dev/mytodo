// Package commands implements Command pattern for comments write operations.

package commands

// CreateCommentCommand creates a new comment.
//
// Example:
//   cmd := CreateCommentCommand{
//       EntityID:   "issue-123",
//       EntityType: "issue",
//       AuthorID:   "user-1",
//       Content:    "Great work!",
//   }
