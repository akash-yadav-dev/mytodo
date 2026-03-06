// Package http provides HTTP/REST API endpoints for users module.

package http

// Routes registers all user-related HTTP routes.
//
// In production, route files typically:
// - Group related routes
// - Apply middleware (auth, logging, rate limiting)
// - Register handlers for each endpoint
// - Define API versioning
//
// Example route registration:
//   func RegisterUserRoutes(router *mux.Router, controller *UserController) {
//       api := router.PathPrefix("/api/v1/users").Subrouter()
//       api.Use(authMiddleware, loggingMiddleware)
//
//       api.HandleFunc("", controller.List).Methods("GET")
//       api.HandleFunc("/{id}", controller.Get).Methods("GET")
//       api.HandleFunc("", controller.Create).Methods("POST")
//       api.HandleFunc("/{id}", controller.Update).Methods("PUT")
//       api.HandleFunc("/{id}", controller.Delete).Methods("DELETE")
//   }
