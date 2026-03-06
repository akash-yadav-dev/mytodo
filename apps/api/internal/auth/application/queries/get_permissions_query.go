// Package queries implements Query pattern for read operations.
//
// This file defines the query for retrieving user permissions.

package queries

// GetPermissionsQuery represents a request to retrieve user permissions.
// In production applications, permission queries typically include:
// - User ID to get permissions for
// - Optional: resource type (filter by resource)
// - Optional: include role information
// - Optional: include inherited permissions
// - Optional: context (org, project) for scoped permissions
// - Flatten role hierarchy (resolve inherited permissions)
