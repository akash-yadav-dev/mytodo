// Package handlers orchestrates use cases for the users module.

package handlers

// UserQueryHandler processes user read operations.
//
// In production applications, query handlers typically:
// - Execute queries against repositories
// - Apply business logic for data access
// - Transform entities to DTOs
// - Handle caching (read-through pattern)
// - Implement efficient data fetching
//
// Example structure:
//   type UserQueryHandler struct {
//       userRepo    UserRepository
//       profileRepo ProfileRepository
//       cache       Cache
//   }
//
// Example methods:
//   func (h *UserQueryHandler) HandleGetByID(query GetUserByIDQuery) (*UserDTO, error)
//   func (h *UserQueryHandler) HandleList(query ListUsersQuery) (*PaginatedUsers, error)
//   func (h *UserQueryHandler) HandleSearch(query SearchUsersQuery) ([]UserDTO, error)
