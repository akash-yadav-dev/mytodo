// Package service contains domain services for the projects module.

package service

// ProjectConfigService handles project configuration management.
//
// Example interface:
//   type ProjectConfigService interface {
//       GetConfig(projectID string) (*ProjectConfig, error)
//       UpdateConfig(projectID string, config ProjectConfig) error
//       ResetToDefaults(projectID string) error
//   }
//
// Example usage:
//   config, _ := svc.GetConfig("proj-1")
//   // Returns: &ProjectConfig{DefaultIssueType: "task",...}
