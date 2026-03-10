// Package http provides HTTP/REST API endpoints for organization member management.

package http

import (
	"mytodo/apps/api/pkg/middleware"
	"mytodo/apps/api/pkg/security"

	"github.com/gin-gonic/gin"
)

// RegisterMemberRoutes registers member management routes
//
// API endpoints (all protected, require authentication):
//
//	GET    /api/v1/organizations/:id/members               - List organization members
//	POST   /api/v1/organizations/:id/members               - Add member to organization
//	DELETE /api/v1/organizations/:id/members/:userId       - Remove member from organization
//	PATCH  /api/v1/organizations/:id/members/:userId/role  - Update member role
func RegisterMemberRoutes(router *gin.RouterGroup, controller *MemberController, jwtService *security.JWTService) {
	authMiddleware := middleware.AuthMiddleware(jwtService)

	organizations := router.Group("/organizations")
	organizations.Use(authMiddleware)
	{
		// Member management routes
		organizations.GET("/:id/members", controller.ListMembers)                     // List members
		organizations.POST("/:id/members", controller.AddMember)                      // Add member
		organizations.DELETE("/:id/members/:userId", controller.RemoveMember)         // Remove member
		organizations.PATCH("/:id/members/:userId/role", controller.UpdateMemberRole) // Update member role
	}
}
