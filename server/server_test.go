package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/cacastel/go-jsonplaceholder-middle-end/api/post/domain"
	"github.com/stretchr/testify/assert"
)

func TestPostsRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/posts/?limit=1", nil)
	router.ServeHTTP(w, req)

	posts := []*domain.Post{}
	err := json.Unmarshal(w.Body.Bytes(), &posts)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(posts))

	post := posts[0]

	assert.Equal(t, "*domain.Post", reflect.TypeOf(post).String())
	assert.Equal(t, 1, post.Id)
	assert.Equal(t, 1, post.UserId)
	assert.Equal(t, "string", reflect.TypeOf(post.Title).String())
	assert.Equal(t, "string", reflect.TypeOf(post.Body).String())
}
