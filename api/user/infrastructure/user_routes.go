package infrastructure

import "github.com/gin-gonic/gin"

func SetupUserRoutes(r *gin.RouterGroup) {
	handler := NewUserHandler()

	// /api/users
	router := r.Group("/users")

	// /api/users/
	router.GET("/", handler.GetAllUsers)

	// /api/users/user
	router.GET("/user", handler.GetUser)
	router.POST("/user", handler.SaveUser)
}
