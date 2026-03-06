// Package persistence provides concrete implementations of domain repository interfaces.
//
// Infrastructure/Persistence layer in Clean Architecture:
// - Implements repository interfaces defined in the domain layer
// - Contains database-specific code (SQL, ORM, NoSQL queries)
// - Handles data mapping between database models and domain entities
// - Manages database connections and transactions
//
// In production applications, this folder typically contains:
// - Repository implementations using specific databases (PostgreSQL, MySQL, MongoDB)
// - ORM model definitions (GORM, sqlx, or raw SQL)
// - Data mappers (convert DB models to domain entities and vice versa)
// - Database query builders
// - Migration-related utility functions
// - Connection pooling configuration
//
// Best practices:
// - Isolate database technology from domain layer
// - Handle database errors and convert to domain errors
// - Implement efficient queries (indexes, joins)
// - Use prepared statements to prevent SQL injection
// - Handle NULL values appropriately
// - Implement proper transaction management

package persistence

// UserRepositoryImpl is the PostgreSQL implementation of UserRepository interface.
// In production applications, repository implementations typically:
// - Use connection pooling for performance
// - Implement all interface methods defined in domain/repository
// - Map between database rows/documents and domain entities
// - Handle database-specific errors (unique constraints, foreign keys)
// - Implement optimized queries with proper indexes
// - Cache frequently accessed data
// - Use transactions for multi-step operations
// - Log slow queries for performance monitoring
