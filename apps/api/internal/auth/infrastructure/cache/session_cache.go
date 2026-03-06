// Package cache provides caching implementations for the auth module.
//
// Cache layer in Clean Architecture:
// - Provides fast access to frequently used data
// - Reduces database load
// - Improves application performance
// - Can be transparent to application layer
//
// In production applications, this folder typically contains:
// - Redis/Memcached client wrappers
// - Cache key naming strategies
// - TTL (time-to-live) management
// - Cache invalidation logic
// - Serialization/deserialization helpers
// - Cache warming strategies
// - Cache-aside, write-through patterns
//
// Best practices:
// - Set appropriate TTLs to balance freshness and performance
// - Handle cache failures gracefully (fallback to database)
// - Implement cache invalidation on data updates
// - Monitor cache hit rates
// - Use consistent key naming conventions
// - Compress large cached objects

package cache

// SessionCache manages session data in Redis for fast access.
// In production applications, session caches typically:
// - Store active session data with automatic expiration
// - Implement sliding expiration (extend on activity)
// - Support session invalidation (logout, password change)
// - Handle distributed sessions (multiple app servers)
// - Serialize session data efficiently
// - Implement fallback to database on cache miss
// - Monitor cache size and eviction policies
