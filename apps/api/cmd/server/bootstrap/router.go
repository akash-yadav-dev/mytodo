package bootstrap

import (
	"net/http"

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

	return router
}
