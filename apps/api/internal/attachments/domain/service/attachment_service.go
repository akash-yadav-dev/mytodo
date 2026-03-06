// Package service contains domain services for the attachments module.

package service

// AttachmentService handles attachment business logic.
//
// Example interface:
//   type AttachmentService interface {
//       UploadAttachment(file io.Reader, filename, entityID, entityType, uploaderID string) (*Attachment, error)
//       GetAttachment(attachmentID string) (*Attachment, error)
//       DeleteAttachment(attachmentID string) error
//       ValidateFileType(contentType string) error
//       ValidateFileSize(size int64) error
//   }
//
// Example usage:
//   attachment, _ := svc.UploadAttachment(file, "screenshot.png", "issue-123", "issue", "user-1")
//   // Returns: &Attachment{ID: "att-1", Filename: "screenshot.png", URL: "...",...}
