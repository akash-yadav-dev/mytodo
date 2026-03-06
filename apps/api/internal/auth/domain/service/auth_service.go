// Package service contains domain services that implement business logic
// spanning multiple entities or requiring coordination between entities.
//
// Domain services in Clean Architecture represent business operations that:
// - Don't naturally belong to a single entity
// - Coordinate operations across multiple entities
// - Implement complex business rules and workflows
// - Remain independent of infrastructure concerns
//
// In production applications, this folder typically contains:
// - Business logic that requires multiple repositories
// - Domain calculations and validations
// - Transaction coordination (at domain level)
// - Domain event publishing logic
// - Business workflow orchestration

package service

// AuthService handles core authentication business logic.
// In production applications, auth services typically implement:
// - User registration with validation
// - Login authentication (password verification, account checks)
// - Token generation and validation (JWT, refresh tokens)
// - Password reset workflows
// - Email/phone verification
// - Account activation/deactivation
// - Multi-factor authentication (MFA) setup and verification
// - OAuth integration and social login
// - Brute force protection and rate limiting
// - Audit logging of authentication events
