// Package service contains domain services for the projects module.

package service

// ProjectService handles project business logic.
//
// In production, project services typically implement:
// - Project creation with validation
// - Membership management
// - Access control checks
// - Project archiving/deletion workflows
// - Project key uniqueness validation
//
// Example interface:
//   type ProjectService interface {
//       CreateProject(name, key, ownerID string) (*Project, error)
//       UpdateProject(id string, updates ProjectUpdates) (*Project, error)
//       ArchiveProject(id string) error
//       AddMember(projectID, userID, role string) error
//       RemoveMember(projectID, userID string) error
//       GetMembers(projectID string) ([]ProjectMember, error)
//       CanUserAccess(projectID, userID string) (bool, error)
//   }
//
// Example usage:
//   project, err := svc.CreateProject("My Todo App", "MYTODO", "user-123")
//   // Returns: &Project{ID: "proj-1", Key: "MYTODO", Name: "My Todo App",...}, nil
//
//   err := svc.AddMember("proj-1", "user-456", "member")
//   // Adds user-456 as a member to the project
