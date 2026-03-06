// Package persistence provides concrete repository implementations.

package persistence

// ProfileRepositoryImpl implements ProfileRepository interface.
//
// Example method implementation:
//   func (r *ProfileRepositoryImpl) FindByUserID(userID string) (*entity.Profile, error) {
//       var profile entity.Profile
//       query := "SELECT * FROM profiles WHERE user_id = $1"
//       err := r.db.QueryRow(query, userID).Scan(&profile)
//       return &profile, err
//   }
