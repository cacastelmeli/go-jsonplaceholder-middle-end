package handler

import (
	"net/http"

	postCtrl "github.com/cacastel/go-jsonplaceholder-middle-end/controller/post"
	"github.com/gin-gonic/gin"
)

func GetAllPostsHandler(c *gin.Context) {
	posts, err := postCtrl.GetAllPosts(c.Query("limit"))

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
