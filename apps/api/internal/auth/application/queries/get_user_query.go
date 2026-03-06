// Package queries implements Query pattern for read operations in the application layer.
//
// Queries in CQRS (Command Query Responsibility Segregation) represent:
// - Read operations that don't modify state
// - Data retrieval with filtering, sorting, pagination
// - View models and projections for specific use cases
//
// In production applications, this folder typically contains:
// - Query DTOs with filter/search parameters
// - Query handlers that fetch and transform data
// - Pagination and sorting logic
// - Data projection and mapping
// - Read-optimized queries (can use different data sources)
//
// Best practices:
// - Keep queries side-effect free
// - Return view-specific DTOs, not domain entities
// - Implement efficient pagination
// - Consider caching for frequently accessed data
// - Separate read models for complex reporting

package queries

// GetUserQuery represents a request to retrieve a single user.
// In production applications, get user queries typically include:
// - User ID to retrieve
// - Optional: fields to include/exclude (field selection)
// - Optional: include related data (roles, permissions)
// - Requesting user context (for authorization checks)
