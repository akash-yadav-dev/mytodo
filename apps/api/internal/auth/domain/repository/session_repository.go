// Package repository defines data access interfaces for the auth domain.
//
// This file defines the contract for session data persistence.

package repository

// SessionRepository defines data access methods for session entities.
// In production applications, session repositories typically provide:
// - CreateSession(session) - persist new session
// - FindByID(sessionID) - retrieve session by ID
// - FindByUserID(userID) - get all sessions for a user
// - FindActiveByUserID(userID) - get only active sessions
// - UpdateActivity(sessionID, timestamp) - update last activity
// - InvalidateSession(sessionID) - mark session as logged out
// - InvalidateAllUserSessions(userID) - force logout all user sessions
// - DeleteExpiredSessions() - cleanup expired sessions
// - CountActiveByUserID(userID) - count concurrent sessions
// - FindByToken(token) - lookup session by token
// - UpdateMetadata(sessionID, metadata) - update session data
