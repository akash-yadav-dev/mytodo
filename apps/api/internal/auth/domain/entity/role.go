// Package entity defines the core domain entities for the auth module.
//
// This file contains role entities that represent user roles in the RBAC
// (Role-Based Access Control) authorization system.

package entity

// Role represents a named collection of permissions in the authorization system.
// In production applications, role entities typically include:
// - Unique role ID
// - Role name (e.g., "admin", "user", "manager", "viewer")
// - Display name and description
// - Associated permissions (list of permission IDs)
// - Role hierarchy/level (for role inheritance)
// - Role scope (system-wide, organization-level, project-level)
// - Active/inactive status
// - System role flag (prevents deletion of built-in roles)
// - Timestamps (created_at, updated_at)
// - Role assignment rules and conditions
