// Package dto defines Data Transfer Objects for projects module.

package dto

// Project DTOs represent project data for API requests/responses.
//
// Example structures:
//   type ProjectDTO struct {
//       ID          string           `json:"id"`
//       Key         string           `json:"key"`
//       Name        string           `json:"name"`
//       Description string           `json:"description"`
//       OwnerID     string           `json:"owner_id"`
//       Status      string           `json:"status"`
//       Visibility  string           `json:"visibility"`
//       Members     []ProjectMemberDTO `json:"members,omitempty"`
//       CreatedAt   time.Time        `json:"created_at"`
//       UpdatedAt   time.Time        `json:"updated_at"`
//   }
//
//   type ProjectMemberDTO struct {
//       UserID   string `json:"user_id"`
//       Username string `json:"username"`
//       Role     string `json:"role"`
//       JoinedAt time.Time `json:"joined_at"`
//   }
