// Package service contains domain services for the auth module.
//
// This file implements session management business logic.

package service

// SessionService handles user session lifecycle and management.
// In production applications, session services typically implement:
// - Session creation on successful login
// - Session validation and refresh
// - Session termination (logout)
// - Idle timeout detection and enforcement
// - Concurrent session management (limits, single session enforcement)
// - Session device tracking and management
// - Force logout (admin action, password change)
// - Session activity tracking
// - "Remember me" functionality
// - Session hijacking prevention
