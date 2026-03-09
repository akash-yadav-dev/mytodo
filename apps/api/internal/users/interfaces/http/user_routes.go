// Package http provides HTTP/REST API endpoints for users module.

package http

import (
	"mytodo/apps/api/pkg/middleware"
	"mytodo/apps/api/pkg/security"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registers all user-related HTTP routes.
//
// API endpoints:
//
//	Public:
//	  GET    /api/v1/users           - List all users (paginated)
//	  GET    /api/v1/users/:id       - Get user profile by ID
//	  GET    /api/v1/users/search    - Search users
//
//	Protected (requires authentication):
//	  GET    /api/v1/users/me                - Get current user's profile
//	  POST   /api/v1/users/profile           - Create user profile
//	  PUT    /api/v1/users/me                - Update current user's profile
//	  DELETE /api/v1/users/me                - Delete current user's profile
//	  GET    /api/v1/users/me/preferences    - Get user preferences
//	  PUT    /api/v1/users/me/preferences    - Update user preferences
func RegisterUserRoutes(router *gin.RouterGroup, controller *UserController, jwtService *security.JWTService) {
	users := router.Group("/users")

	// Public routes - no authentication required
	users.GET("", controller.ListUserProfiles)          // List users (paginated)
	users.GET("/search", controller.SearchUserProfiles) // Search users
	users.GET("/:id", controller.GetUserProfileByID)    // Get user profile by ID

	// Protected routes - authentication required
	authMiddleware := middleware.AuthMiddleware(jwtService)
	protected := users.Group("")
	protected.Use(authMiddleware)
	{
		// Current user profile management
		protected.GET("/me", controller.GetCurrentUserProfile)   // Get own profile
		protected.POST("/profile", controller.CreateUserProfile) // Create profile (after registration)
		protected.PUT("/me", controller.UpdateUserProfile)       // Update own profile
		protected.DELETE("/me", controller.DeleteUserProfile)    // Delete own profile

		// User preferences management
		protected.GET("/me/preferences", controller.GetUserPreferences)    // Get preferences
		protected.PUT("/me/preferences", controller.UpdateUserPreferences) // Update preferences
	}
}
