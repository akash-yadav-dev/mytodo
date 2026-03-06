// Package entity defines core domain entities for the attachments module.
//
// Attachments represent files uploaded to issues, comments, etc.

package entity

// Attachment represents a file attached to an entity.
//
// In production, Attachment entities include:
// - ID, Filename, ContentType, Size
// - URL/Path (storage location)
// - EntityID, EntityType
// - UploaderID, UploadedAt
// - Metadata (dimensions for images, duration for videos)
//
// Example structure:
//   type Attachment struct {
//       ID          string    `json:"id"`
//       Filename    string    `json:"filename"`
//       ContentType string    `json:"content_type"`
//       Size        int64     `json:"size"`
//       URL         string    `json:"url"`
//       EntityID    string    `json:"entity_id"`
//       EntityType  string    `json:"entity_type"`
//       UploaderID  string    `json:"uploader_id"`
//       UploadedAt  time.Time `json:"uploaded_at"`
//   }
