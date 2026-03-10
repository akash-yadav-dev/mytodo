// Package http provides HTTP/REST API endpoints for the organizations module.
//
// This file defines route registration and grouping.

package http

import (
	"mytodo/apps/api/pkg/middleware"
	"mytodo/apps/api/pkg/security"

	"github.com/gin-gonic/gin"
)

// RegisterOrganizationRoutes registers all organization-related routes
//
// API endpoints:
//
//	Public:
//	  GET    /api/v1/organizations           - List all organizations (paginated)
//	  GET    /api/v1/organizations/search    - Search organizations
//	  GET    /api/v1/organizations/:id       - Get organization by ID
//
//	Protected (requires authentication):
//	  POST   /api/v1/organizations                 - Create new organization
//	  PUT    /api/v1/organizations/:id             - Update organization
//	  DELETE /api/v1/organizations/:id             - Delete organization (soft delete)
//	  GET    /api/v1/organizations/me/owned        - Get my owned organizations
//	  GET    /api/v1/organizations/me/member       - Get organizations where I'm a member
//	  POST   /api/v1/organizations/:id/transfer    - Transfer organization ownership
func RegisterOrganizationRoutes(router *gin.RouterGroup, controller *OrganizationController, jwtService *security.JWTService) {
	organizations := router.Group("/organizations")

	// Public routes - no authentication required
	organizations.GET("", controller.ListOrganizations)          // List organizations (paginated)
	organizations.GET("/search", controller.SearchOrganizations) // Search organizations
	organizations.GET("/:id", controller.GetOrganization)        // Get organization by ID

	// Protected routes - authentication required
	authMiddleware := middleware.AuthMiddleware(jwtService)
	protected := organizations.Group("")
	protected.Use(authMiddleware)
	{
		// Organization CRUD operations
		protected.POST("", controller.CreateOrganization)       // Create organization
		protected.PUT("/:id", controller.UpdateOrganization)    // Update organization
		protected.DELETE("/:id", controller.DeleteOrganization) // Delete organization

		// User's organizations
		protected.GET("/me/owned", controller.GetMyOrganizations)      // Get owned organizations
		protected.GET("/me/member", controller.GetMemberOrganizations) // Get member organizations

		// Ownership transfer
		protected.POST("/:id/transfer", controller.TransferOwnership) // Transfer ownership
	}
}
