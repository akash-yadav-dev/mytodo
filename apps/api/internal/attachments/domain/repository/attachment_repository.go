// Package repository defines data access interfaces for attachments.

package repository

// AttachmentRepository defines data access methods for attachments.
//
// Example interface:
//   type AttachmentRepository interface {
//       Create(attachment *Attachment) error
//       FindByID(id string) (*Attachment, error)
//       FindByEntity(entityID string) ([]Attachment, error)
//       Delete(id string) error
//   }
//
// Example usage:
//   attachments, _ := repo.FindByEntity("issue-123")
//   // Returns: []Attachment{{Filename: "screenshot.png"},...}
