// Package entity defines core domain entities for the projects module.

package entity

// ProjectMember represents a user's membership in a project.
//
// In production applications, ProjectMember entities include:
// - ProjectID, UserID
// - Role (owner, admin, member, viewer)
// - Permissions (custom permission overrides)
// - JoinedAt timestamp
//
// Example structure:
//   type ProjectMember struct {
//       ProjectID string   `json:"project_id"`
//       UserID    string   `json:"user_id"`
//       Role      string   `json:"role"`
//       JoinedAt  time.Time `json:"joined_at"`
//   }
//
// Example methods:
//   func (pm *ProjectMember) HasPermission(permission string) bool
//   func (pm *ProjectMember) IsOwner() bool
