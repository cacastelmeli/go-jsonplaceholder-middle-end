package infrastructure

import (
	"github.com/cacastel/go-jsonplaceholder-middle-end/api/post/domain"
	"github.com/cacastel/go-jsonplaceholder-middle-end/api/shared/infrastructure"
)

type JpPostRepository struct {
	client *infrastructure.JpClient
}

func NewJpPostRepository() *JpPostRepository {
	return &JpPostRepository{
		client: infrastructure.NewJpClient(),
	}
}

func (j *JpPostRepository) FetchAllPosts() ([]*domain.Post, error) {
	return j.client.GetAllPosts()
}
