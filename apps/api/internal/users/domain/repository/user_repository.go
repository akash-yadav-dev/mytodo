// Package repository defines data access interfaces for the users domain.
//
// Repository interfaces define contracts for data persistence without
// specifying implementation details (database-agnostic).

package repository

// UserRepository defines data access methods for user entities.
//
// In production applications, user repositories typically provide:
// - CRUD operations (Create, Read, Update, Delete)
// - Query methods (FindByID, FindByEmail, FindByUsername)
// - Search and filtering with pagination
// - Existence checks for uniqueness validation
// - Bulk operations for efficiency
//
// Example interface:
//   type UserRepository interface {
//       Create(user *User) error
//       FindByID(id string) (*User, error)
//       FindByEmail(email string) (*User, error)
//       FindByUsername(username string) (*User, error)
//       Update(user *User) error
//       Delete(id string) error
//       List(filters Filters, pagination Pagination) ([]User, int, error)
//       ExistsByEmail(email string) (bool, error)
//       ExistsByUsername(username string) (bool, error)
//       Search(query string, limit int) ([]User, error)
//   }
//
// Example usage:
//   user, err := repo.FindByEmail("user@example.com")
//   // Returns: &User{ID: "123", Email: "user@example.com", ...}, nil
//   // Returns: nil, ErrUserNotFound if not found
//
//   exists, _ := repo.ExistsByUsername("johndoe")
//   // Returns: true if username exists, false otherwise
//
//   users, total, err := repo.List(Filters{Status: "active"}, Pagination{Page: 1, Limit: 20})
//   // Returns: []User{...}, 150 (total count), nil
