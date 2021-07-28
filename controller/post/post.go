package post

import (
	"errors"
	"strconv"

	"github.com/cacastel/go-jsonplaceholder-middle-end/client"
)

var jsonPlaceholderClient = client.NewJsonPlaceholderClient()

var errLimit = errors.New("limit out of bounds")

func GetAllPosts(limit string) ([]*client.JsonPlaceholderPost, error) {
	posts, err := jsonPlaceholderClient.GetAllPosts()

	if err != nil {
		// c.AbortWithStatus(http.StatusInternalServerError)
		return nil, err
	}

	// limit := c.Query("limit")

	if len(limit) != 0 {
		iLimit, err := strconv.Atoi(limit)

		if err != nil {
			return nil, err
			// c.AbortWithStatus(http.StatusBadRequest)
		}

		if iLimit <= 0 {
			return nil, errLimit
		}

		posts = posts[:iLimit]
	}

	// c.JSON(http.StatusOK, gin.H{"posts": posts})
	return posts, nil
}
