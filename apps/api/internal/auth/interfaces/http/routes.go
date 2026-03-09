// Package http provides HTTP/REST API endpoints for the auth module.
//
// This file defines route registration and grouping.

package http

import (
	"mytodo/apps/api/pkg/middleware"
	"mytodo/apps/api/pkg/security"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes registers all auth-related routes
func RegisterAuthRoutes(router *gin.RouterGroup, controller *AuthController, jwtService *security.JWTService) {
	auth := router.Group("/auth")
	{
		// Public routes
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
		auth.POST("/refresh", controller.RefreshToken)
		auth.POST("/logout", controller.Logout)

		// Protected routes
		auth.GET("/me", middleware.AuthMiddleware(jwtService), controller.Me)
	}
}
