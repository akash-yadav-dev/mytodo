// Package entity defines core domain entities for the projects module.

package entity

// ProjectConfig represents project-specific configuration settings.
//
// In production, project configs typically include:
// - Default issue settings
// - Workflow configurations
// - Notification preferences
// - Integration settings
// - Custom fields
//
// Example structure:
//   type ProjectConfig struct {
//       ProjectID             string            `json:"project_id"`
//       DefaultIssueType      string            `json:"default_issue_type"`
//       AllowedIssueTypes     []string          `json:"allowed_issue_types"`
//       EnableNotifications   bool              `json:"enable_notifications"`
//       CustomFields          map[string]string `json:"custom_fields"`
//   }
