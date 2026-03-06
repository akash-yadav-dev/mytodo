// Package service contains domain services for the boards module.

package service

// BoardService handles board business logic.
//
// Example interface:
//   type BoardService interface {
//       CreateBoard(projectID, name string) (*Board, error)
//       AddColumn(boardID, name string, position int) error
//       MoveCard(cardID, toColumnID string, position int) error
//       ValidateWIPLimit(columnID string) error
//   }
//
// Example usage:
//   board, _ := svc.CreateBoard("proj-1", "Sprint Board")
//   // Returns: &Board{ID: "board-1", Name: "Sprint Board",...}
