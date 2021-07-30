package infrastructure

import "github.com/gin-gonic/gin"

func SetupPostRoutes(r *gin.RouterGroup) {
	handler := NewPostHandler()

	router := r.Group("/posts")
	router.GET("/", handler.GetPostsHandler)
}
