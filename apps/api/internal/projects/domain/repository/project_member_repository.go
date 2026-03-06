// Package repository defines data access interfaces for the projects domain.

package repository

// ProjectMemberRepository manages project membership data.
//
// Example interface:
//   type ProjectMemberRepository interface {
//       Add(member *ProjectMember) error
//       Remove(projectID, userID string) error
//       FindByProject(projectID string) ([]ProjectMember, error)
//       FindByUser(userID string) ([]ProjectMember, error)
//       IsMember(projectID, userID string) (bool, error)
//       UpdateRole(projectID, userID, newRole string) error
//   }
//
// Example usage:
//   members, _ := repo.FindByProject("proj-1")
//   // Returns: []ProjectMember{{UserID: "user-1", Role: "owner"},...}
//
//   isMember, _ := repo.IsMember("proj-1", "user-123")
//   // Returns: true if user is a member, false otherwise
