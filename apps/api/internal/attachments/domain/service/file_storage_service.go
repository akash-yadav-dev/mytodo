// Package service contains domain services for the attachments module.

package service

// FileStorageService handles file storage operations (S3, local, etc.).
//
// Example interface:
//   type FileStorageService interface {
//       Upload(file io.Reader, key string) (string, error)
//       Download(key string) (io.ReadCloser, error)
//       Delete(key string) error
//       GetURL(key string) (string, error)
//   }
//
// Example usage:
//   url, _ := storageSvc.Upload(file, "attachments/issue-123/screenshot.png")
//   // Returns: "https://cdn.example.com/attachments/issue-123/screenshot.png"
