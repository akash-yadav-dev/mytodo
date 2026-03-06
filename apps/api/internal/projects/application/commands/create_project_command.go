// Package commands implements Command pattern for project write operations.

package commands

// CreateProjectCommand creates a new project.
//
// Example structure:
//   type CreateProjectCommand struct {
//       Name        string `json:"name" validate:"required,min=3,max=100"`
//       Key         string `json:"key" validate:"required,uppercase,min=2,max=10"`
//       Description string `json:"description,omitempty"`
//       Visibility  string `json:"visibility" validate:"oneof=public private internal"`
//       OwnerID     string `json:"owner_id" validate:"required"`
//   }
//
// Example usage:
//   cmd := CreateProjectCommand{
//       Name:       "My Todo Application",
//       Key:        "MYTODO",
//       Visibility: "private",
//       OwnerID:    "user-123",
//   }
//   result, err := handler.Handle(cmd)
//   // Returns: &ProjectDTO{ID: "proj-1", Key: "MYTODO",...}, nil
