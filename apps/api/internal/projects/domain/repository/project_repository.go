// Package repository defines data access interfaces for the projects domain.

package repository

// ProjectRepository defines data access methods for projects.
//
// Example interface:
//   type ProjectRepository interface {
//       Create(project *Project) error
//       FindByID(id string) (*Project, error)
//       FindByKey(key string) (*Project, error)
//       Update(project *Project) error
//       Delete(id string) error
//       List(userID string, filters Filters) ([]Project, error)
//       ExistsByKey(key string) (bool, error)
//   }
//
// Example usage:
//   project, err := repo.FindByKey("MYTODO")
//   // Returns: &Project{ID: "proj-1", Key: "MYTODO",...}, nil
//
//   exists, _ := repo.ExistsByKey("NEWPROJ")
//   // Returns: true if key exists, false otherwise
