// Package service contains domain services for the users module.
//
// Domain services implement business logic that doesn't naturally fit
// within a single entity and coordinates operations across entities.

package service

// UserService handles user management business logic.
//
// In production applications, user services typically implement:
// - User creation with validation
// - Profile updates with business rules
// - User search and filtering
// - Username/email uniqueness validation
// - User deactivation workflows
// - Bulk user operations
//
// Example interface:
//   type UserService interface {
//       CreateUser(email, username, name string) (*User, error)
//       UpdateUser(id string, updates UserUpdates) (*User, error)
//       GetUser(id string) (*User, error)
//       SearchUsers(query string, filters Filters) ([]User, error)
//       DeactivateUser(id string) error
//       ValidateUsername(username string) error
//   }
//
// Example usage:
//   user, err := userService.CreateUser("user@example.com", "johndoe", "John Doe")
//   // Returns: &User{ID: "uuid-123", Email: "user@example.com", ...}, nil
//
//   err := userService.ValidateUsername("johndoe")
//   // Returns: ErrUsernameTaken if already exists, nil otherwise
