package infrastructure

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/cacastel/go-jsonplaceholder-middle-end/api/user/application"
	"github.com/cacastel/go-jsonplaceholder-middle-end/api/user/domain"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	searcher *application.UserSearcher
	saver    *application.UserSaver
}

var (
	ErrMissingIdField = errors.New("the 'id' field is missing")
	ErrIdWrongType    = errors.New("the 'id' has a wrong a type")
)

func NewUserHandler() *UserHandler {
	repo := &FakeUserRepository{
		users: []*domain.User{},
	}

	return &UserHandler{
		searcher: &application.UserSearcher{
			Repo: repo,
		},
		saver: &application.UserSaver{
			Repo: repo,
		},
	}
}

func (u *UserHandler) GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, u.searcher.SearchAll())
}

func (u *UserHandler) GetUser(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.AbortWithError(http.StatusBadRequest, ErrMissingIdField)
		return
	}

	iId, err := strconv.Atoi(id)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, ErrIdWrongType)
		return
	}

	c.JSON(http.StatusOK, u.searcher.Search(iId))
}

func (u *UserHandler) SaveUser(c *gin.Context) {
	userToSave := &domain.User{}

	c.BindJSON(&userToSave)
	u.saver.Save(userToSave)

	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
