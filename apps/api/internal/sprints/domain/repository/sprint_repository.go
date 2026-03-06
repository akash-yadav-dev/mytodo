// Package repository defines data access interfaces for sprints.

package repository

// SprintRepository defines data access methods for sprints.
//
// Example interface:
//   type SprintRepository interface {
//       Create(sprint *Sprint) error
//       FindByID(id string) (*Sprint, error)
//       FindByProject(projectID string) ([]Sprint, error)
//       FindActiveSprint(projectID string) (*Sprint, error)
//       Update(sprint *Sprint) error
//   }
