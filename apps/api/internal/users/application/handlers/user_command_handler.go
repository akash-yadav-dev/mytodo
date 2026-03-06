// Package handlers orchestrates use cases for the users module.
//
// Handlers coordinate domain services and repositories to execute commands.

package handlers

// UserCommandHandler processes user write operations.
//
// In production applications, command handlers typically:
// - Validate command data
// - Coordinate with domain services
// - Manage transactions
// - Emit domain events
// - Return DTOs (not domain entities)
//
// Example structure:
//   type UserCommandHandler struct {
//       userService UserService
//       userRepo    UserRepository
//       eventBus    EventBus
//   }
//
// Example methods:
//   func (h *UserCommandHandler) HandleCreate(cmd CreateUserCommand) (*UserDTO, error)
//   func (h *UserCommandHandler) HandleUpdate(cmd UpdateUserCommand) (*UserDTO, error)
//   func (h *UserCommandHandler) HandleDelete(cmd DeleteUserCommand) error
