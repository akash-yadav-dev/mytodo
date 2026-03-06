// Package commands implements Command pattern for attachments write operations.

package commands

// UploadAttachmentCommand uploads a file attachment.
//
// Example:
//   cmd := UploadAttachmentCommand{
//       File:       fileReader,
//       Filename:   "screenshot.png",
//       EntityID:   "issue-123",
//       EntityType: "issue",
//       UploaderID: "user-1",
//   }
