// Package handlers contains application-layer handlers that orchestrate use cases.
//
// Handlers in Clean Architecture:
// - Orchestrate business workflows by coordinating domain services
// - Handle command and query execution
// - Manage transactions and unit of work
// - Convert between DTOs and domain entities
// - Handle application-level concerns (logging, validation, events)
//
// In production applications, this folder typically contains:
// - Command handlers (process write operations)
// - Query handlers (process read operations)
// - Use case implementations
// - Transaction management
// - Domain event dispatching
// - Application-level validation
//
// Best practices:
// - Keep handlers thin - delegate to domain services
// - Handle cross-cutting concerns (auth, logging, metrics)
// - Manage transactional boundaries
// - Return application DTOs, not domain entities
// - Handle errors and map to application-level error codes

package handlers

// AuthHandler orchestrates authentication-related use cases.
// In production applications, auth handlers typically implement:
// - Handle(LoginCommand) - process login requests
// - Handle(RegisterCommand) - process registrations
// - Handle(LogoutCommand) - handle logout
// - Handle(RefreshTokenCommand) - token refresh
// - Handle(ResetPasswordCommand) - password reset
// - Coordinate domain services (auth, token, session services)
// - Manage database transactions
// - Emit domain events (UserLoggedIn, UserRegistered)
// - Handle validation errors and authentication failures
