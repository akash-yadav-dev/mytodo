// Package queries implements Query pattern for read operations.
//
// This file defines the list users query.

package queries

// ListUsersQuery represents a request to retrieve multiple users.
// In production applications, list users queries typically include:
// - Pagination parameters (page, limit/per_page, offset)
// - Sorting parameters (sort_by, order: asc/desc)
// - Filter conditions (status, role, created date range)
// - Search query (search across name, email)
// - Optional: include counts (total records)
// - Optional: cursor-based pagination token
// - Field selection (which columns to return)
