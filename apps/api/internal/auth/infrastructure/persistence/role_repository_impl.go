// Package persistence provides concrete implementations of repository interfaces.
//
// This file implements role and permission repository.

package persistence

// RoleRepositoryImpl implements RoleRepository interface.
// In production applications, role repositories typically:
// - Cache role and permission data (rarely changes)
// - Implement efficient permission checking queries
// - Handle role hierarchy and inheritance
// - Support many-to-many relationships (users-roles, roles-permissions)
// - Prevent deletion of roles still in use
// - Implement bulk role assignment operations
// - Track audit history for security compliance
