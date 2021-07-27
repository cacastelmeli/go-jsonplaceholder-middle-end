package server

import (
	"net/http"
	"os"

	postCtrl "github.com/cacastel/go-jsonplaceholder-middle-end/controller/post"
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
	group := r.Group("/api")

	group.GET("/posts", postCtrl.GetAllPosts)
}

func setupNoRoute(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
}

func StartGinServer() {
	router := gin.Default()

	setupApiRoutes(router)
	setupNoRoute(router)

	router.Run(":" + getServerPort())
}
