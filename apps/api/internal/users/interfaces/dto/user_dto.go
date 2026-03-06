// Package dto defines Data Transfer Objects for users module.
//
// DTOs decouple API contracts from domain entities.

package dto

// User DTOs represent user data for API requests/responses.
//
// Example structures:
//   type UserDTO struct {
//       ID          string    `json:"id"`
//       Email       string    `json:"email"`
//       Username    string    `json:"username"`
//       DisplayName string    `json:"display_name"`
//       Avatar      string    `json:"avatar,omitempty"`
//       Status      string    `json:"status"`
//       CreatedAt   time.Time `json:"created_at"`
//       UpdatedAt   time.Time `json:"updated_at"`
//   }
