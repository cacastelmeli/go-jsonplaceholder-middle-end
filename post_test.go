package main

import (
	"testing"

	"github.com/cacastel/go-jsonplaceholder-middle-end/controller/post"
	"github.com/stretchr/testify/assert"
)

/*
type MockedGinContext struct {
	mock.Mock
	*gin.Context
}

func (m *MockedGinContext) AbortWithStatus(code int) {
	m.Called(code)
}*/

func TestGetAllPosts(t *testing.T) {
	// wrong path
	posts, err := post.GetAllPosts("-4")

	assert.NotNil(t, err)
	assert.Nil(t, posts)

	// happy path
	posts, err = post.GetAllPosts("1")

	assert.Nil(t, err)
	assert.NotNil(t, posts)
	assert.Equal(t, 1, len(posts))
}
