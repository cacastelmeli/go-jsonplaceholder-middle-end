package post

import (
	"net/http"
	"strconv"

	"github.com/cacastel/go-jsonplaceholder-middle-end/client"
	"github.com/gin-gonic/gin"
)

var jsonPlaceholderClient = client.NewJsonPlaceholderClient()

func GetAllPosts(c *gin.Context) {
	posts, err := jsonPlaceholderClient.GetAllPosts()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	limit := c.Query("limit")

	if len(limit) != 0 {
		iLimit, err := strconv.Atoi(limit)

		if err != nil || iLimit <= 0 {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		posts = posts[:iLimit]
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
