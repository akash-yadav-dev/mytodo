// Package http provides HTTP/REST API endpoints for projects module.

package http

// ProjectController handles HTTP endpoints for project operations.
//
// Example endpoints:
//   GET    /api/v1/projects           - List projects
//   GET    /api/v1/projects/:id       - Get project details
//   POST   /api/v1/projects           - Create project
//   PUT    /api/v1/projects/:id       - Update project
//   DELETE /api/v1/projects/:id       - Delete project
//   POST   /api/v1/projects/:id/archive - Archive project
//   GET    /api/v1/projects/:id/members - List project members
//   POST   /api/v1/projects/:id/members - Add member
