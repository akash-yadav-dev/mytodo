// Package queries implements Query pattern for boards read operations.

package queries

// GetBoardQuery retrieves a board with its columns and cards.
//
// Example usage:
//   query := GetBoardQuery{BoardID: "board-1", IncludeCards: true}
//   board, _ := handler.Handle(query)
//   // Returns: &BoardDTO{ID: "board-1", Columns: [...],...}
