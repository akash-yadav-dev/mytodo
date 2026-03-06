// Package http provides HTTP/REST API endpoints for users module.
//
// Controllers handle HTTP requests and delegate to application handlers.

package http

// UserController handles HTTP endpoints for user operations.
//
// In production, HTTP controllers typically:
// - Parse and validate request data
// - Call application handlers
// - Format responses with appropriate status codes
// - Handle errors and return proper error responses
//
// Example endpoints:
//   GET    /api/v1/users           - List users
//   GET    /api/v1/users/:id       - Get user by ID
//   POST   /api/v1/users           - Create user
//   PUT    /api/v1/users/:id       - Update user
//   DELETE /api/v1/users/:id       - Delete user
//   GET    /api/v1/users/search    - Search users
//
// Example method:
//   func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
//       userID := mux.Vars(r)["id"]
//       query := GetUserByIDQuery{UserID: userID}
//       user, err := c.queryHandler.Handle(query)
//       if err != nil {
//           http.Error(w, err.Error(), http.StatusNotFound)
//           return
//       }
//       json.NewEncoder(w).Encode(user)
//   }
