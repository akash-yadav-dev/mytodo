// Package cache provides caching implementations for projects.

package cache

// ProjectCache manages project data caching.
//
// Example methods:
//   func (c *ProjectCache) Get(projectID string) (*Project, error)
//   func (c *ProjectCache) Set(projectID string, project *Project) error
//   func (c *ProjectCache) InvalidateByKey(key string) error
