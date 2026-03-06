// Package commands implements Command pattern for write operations in the application layer.
//
// Commands in CQRS (Command Query Responsibility Segregation) represent:
// - Write operations that modify system state
// - Business use cases that change data
// - Operations that have side effects
//
// In production applications, this folder typically contains:
// - Command DTOs (Data Transfer Objects) with input data
// - Command validation logic
// - Command handlers that execute the business logic
// - Idempotency handling
// - Command metadata (user context, correlation IDs)
//
// Best practices:
// - Keep commands focused and single-purpose
// - Validate command data before processing
// - Return results or errors, not domain entities
// - Handle transactional boundaries
// - Emit domain events after successful execution

package commands

// LoginCommand represents a user login request.
// In production applications, login commands typically include:
// - Email or username
// - Password (plain text, will be hashed for comparison)
// - Optional: device information for session tracking
// - Optional: IP address for security logging
// - Optional: remember me flag
// - Optional: MFA code/token
