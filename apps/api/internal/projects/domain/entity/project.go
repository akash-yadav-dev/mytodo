// Package entity defines core domain entities for the projects module.
//
// Projects are the primary organizational units in the system.

package entity

// Project represents a project in the system.
//
// In production applications, Project entities typically include:
// - ID, Name, Key (unique identifier like "PROJ-123")
// - Description, Owner
// - Status (active, archived, on_hold)
// - Visibility (public, private, internal)
// - Settings and configuration
// - Timestamps (created_at, updated_at, archived_at)
//
// Example structure:
//   type Project struct {
//       ID          string    `json:"id"`
//       Key         string    `json:"key"` // e.g., "MYTODO"
//       Name        string    `json:"name"`
//       Description string    `json:"description"`
//       OwnerID     string    `json:"owner_id"`
//       Status      string    `json:"status"`
//       Visibility  string    `json:"visibility"`
//       CreatedAt   time.Time `json:"created_at"`
//   }
//
// Example methods:
//   func (p *Project) Archive() error
//   func (p *Project) CanAccess(userID string) bool
//   func (p *Project) AddMember(userID string, role string) error
