// Package persistence provides concrete repository implementations for projects.

package persistence

// ProjectRepositoryImpl implements ProjectRepository using PostgreSQL.
//
// Example method:
//   func (r *ProjectRepositoryImpl) FindByKey(key string) (*entity.Project, error) {
//       var project entity.Project
//       err := r.db.Where("key = ?", key).First(&project).Error
//       return &project, err
//   }
