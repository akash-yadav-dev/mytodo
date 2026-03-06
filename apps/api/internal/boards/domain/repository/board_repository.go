// Package repository defines data access interfaces for boards.

package repository

// BoardRepository defines data access methods for boards.
//
// Example interface:
//   type BoardRepository interface {
//       Create(board *Board) error
//       FindByID(id string) (*Board, error)
//       FindByProject(projectID string) ([]Board, error)
//       Update(board *Board) error
//       Delete(id string) error
//   }
//
// Example usage:
//   boards, _ := repo.FindByProject("proj-1")
//   // Returns: []Board{{ID: "board-1", Name: "Main Board"},...}
