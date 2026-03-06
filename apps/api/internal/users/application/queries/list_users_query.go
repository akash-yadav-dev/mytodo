// Package queries implements Query pattern for read operations.

package queries

// ListUsersQuery retrieves a paginated list of users with filters.
//
// Example structure:
//   type ListUsersQuery struct {
//       Page      int               `json:"page" validate:"min=1"`
//       Limit     int               `json:"limit" validate:"min=1,max=100"`
//       SortBy    string            `json:"sort_by"` // "created_at", "email"
//       SortOrder string            `json:"sort_order"` // "asc", "desc"
//       Filters   map[string]string `json:"filters"` // status, role, etc.
//   }
//
// Example usage:
//   query := ListUsersQuery{
//       Page:      1,
//       Limit:     20,
//       SortBy:    "created_at",
//       SortOrder: "desc",
//       Filters:   map[string]string{"status": "active"},
//   }
//   result, err := handler.Handle(query)
//   // Returns: &PaginatedUsers{
//   //   Users: []UserDTO{{ID: "1",...}, {ID: "2",...}},
//   //   Total: 150,
//   //   Page:  1,
//   //   Limit: 20,
//   // }, nil
