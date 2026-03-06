// Package persistence provides concrete repository implementations for users.
//
// This implements UserRepository interface using PostgreSQL/database.

package persistence

// UserRepositoryImpl implements UserRepository using a SQL database.
//
// In production, repository implementations typically:
// - Use ORM (GORM, sqlx) or raw SQL
// - Handle connection pooling
// - Implement efficient queries with indexes
// - Map database models to domain entities
// - Handle database-specific errors
//
// Example structure:
//   type UserRepositoryImpl struct {
//       db *sql.DB
//   }
//
// Example method implementation:
//   func (r *UserRepositoryImpl) FindByID(id string) (*entity.User, error) {
//       var user entity.User
//       err := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user)
//       if err == sql.ErrNoRows {
//           return nil, ErrUserNotFound
//       }
//       return &user, err
//   }
