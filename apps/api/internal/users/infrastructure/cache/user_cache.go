// Package cache provides caching implementations for users module.
//
// Caching layer improves performance by reducing database queries.

package cache

// UserCache manages user data caching using Redis.
//
// In production, cache implementations typically:
// - Use Redis/Memcached for distributed caching
// - Implement cache-aside pattern
// - Set appropriate TTLs
// - Handle cache invalidation
// - Serialize/deserialize data efficiently
//
// Example structure:
//   type UserCache struct {
//       redis *redis.Client
//       ttl   time.Duration
//   }
//
// Example methods:
//   func (c *UserCache) Get(userID string) (*User, error)
//   func (c *UserCache) Set(userID string, user *User) error
//   func (c *UserCache) Delete(userID string) error
//   func (c *UserCache) GetMultiple(userIDs []string) (map[string]*User, error)
//
// Example usage:
//   user, err := cache.Get("user-123")
//   // Returns: cached user if exists, nil + ErrCacheMiss if not
