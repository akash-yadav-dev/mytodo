package bootstrap

import (
	"net/http"

	authRoutes "mytodo/apps/api/internal/auth/interfaces/http"
	orgRoutes "mytodo/apps/api/internal/organizations/interfaces/http"
	userRoutes "mytodo/apps/api/internal/users/interfaces/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(container *Container) *gin.Engine {

	router := gin.New()

	// Production middleware
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.GET("/health", func(c *gin.Context) {
		container.Log.Info("Health check endpoint called")
		c.String(http.StatusOK, "server is running")
	})

	// API v1 routes
	v1 := router.Group("/api/v1")

	authRoutes.RegisterAuthRoutes(
		v1,
		container.AuthController,
		container.JWTService,
	)
	userRoutes.RegisterUserRoutes(
		v1,
		container.UserController,
		container.JWTService,
	)
	orgRoutes.RegisterOrganizationRoutes(
		v1,
		container.OrgController,
		container.JWTService,
	)
	orgRoutes.RegisterMemberRoutes(
		v1,
		container.MemberController,
		container.JWTService,
	)

	return router
}
