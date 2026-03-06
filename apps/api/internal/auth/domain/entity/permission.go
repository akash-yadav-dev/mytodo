// Package entity defines the core domain entities for the auth module.
//
// This file contains permission entities that define granular access rights
// in the system's authorization model.

package entity

// Permission represents a single access right or capability in the system.
// In production applications, permission entities typically include:
// - Unique permission ID
// - Permission name/key (e.g., "issues:create", "projects:delete")
// - Resource type (what entity this permission applies to)
// - Action type (read, write, delete, admin)
// - Display name and description
// - Permission scope (global, organization, project, personal)
// - Category/grouping (for UI organization)
// - Active status
// - System permission flag (for built-in permissions)
// - Derived/computed permissions (permissions implied by this one)
