package server

import (
	"net/http"
	"os"

	"github.com/cacastel/go-jsonplaceholder-middle-end/api/post/infrastructure"
	"github.com/gin-gonic/gin"
)

func getServerPort() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8000"
	}

	return port
}

func setupApiRoutes(r *gin.Engine) {
	rootGroup := r.Group("/api")

	infrastructure.SetupPostRoutes(rootGroup)
}

func setupNoRoute(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	setupApiRoutes(router)
	setupNoRoute(router)

	return router
}

func StartGinServer() {
	router := setupRouter()
	router.Run(":" + getServerPort())
}
