package infrastructure

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/cacastel/go-jsonplaceholder-middle-end/api/post/application"
	"github.com/cacastel/go-jsonplaceholder-middle-end/api/post/domain"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	searcher *application.PostSearcher
}

var (
	ErrInvalidLimit = errors.New("invalid limit")
)

func NewPostHandler() *PostHandler {
	return &PostHandler{
		searcher: &application.PostSearcher{
			Repo: NewJpPostRepository(),
		},
	}
}

func (h *PostHandler) GetPostsHandler(c *gin.Context) {
	var posts []*domain.Post
	var errSearch error

	limit, ok := c.GetQuery("limit")

	if !ok {
		posts, errSearch = h.searcher.SearchAll()
	} else {
		nLimit, err := strconv.Atoi(limit)

		if err != nil || nLimit <= 0 {
			c.AbortWithError(http.StatusBadRequest, ErrInvalidLimit)
			return
		}

		posts, errSearch = h.searcher.Search(uint(nLimit))
	}

	if errSearch != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, posts)
}
