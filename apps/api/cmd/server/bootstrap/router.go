package bootstrap

import (
	"net/http"

	authhttp "mytodo/apps/api/internal/auth/interfaces/http"

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

	authhttp.RegisterAuthRoutes(
		v1,
		container.AuthController,
		container.JWTService,
	)

	return router
}
