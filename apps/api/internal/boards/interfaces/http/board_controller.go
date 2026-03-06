// Package http provides HTTP/REST API endpoints for boards.

package http

// BoardController handles HTTP endpoints for board operations.
//
// Endpoints:
//   GET    /api/v1/projects/:projectId/boards
//   POST   /api/v1/boards
//   PUT    /api/v1/boards/:id
//   POST   /api/v1/boards/cards/:cardId/move
