// Package entity defines the core domain entities (business objects) for the auth module.
//
// Domain entities represent the fundamental business concepts and rules in Clean Architecture.
// In production-grade applications, this folder typically contains:
// - Core business objects with their properties and validation logic
// - Business rules and invariants that must always be true
// - Value objects that are immutable and defined by their attributes
// - Domain events that represent significant business occurrences
// - Entity lifecycle methods (creation, validation, state transitions)
//
// Best practices:
// - Entities should be independent of infrastructure concerns (databases, APIs, frameworks)
// - Keep business logic within entities when it relates to a single entity
// - Use value objects for concepts without identity (email, password hash)
// - Validate invariants in constructors or factory methods
// - Entities should have clear identity (usually an ID)

package entity

// User represents an authenticated user in the system.
// In production applications, user entities typically include:
// - Unique identifier (ID, UUID)
// - Authentication credentials (email, hashed password)
// - Profile information (name, avatar)
// - Account status (active, suspended, deleted)
// - Security metadata (failed login attempts, MFA settings)
// - Timestamps (created_at, updated_at, last_login)
// - Relationships to roles and permissions
