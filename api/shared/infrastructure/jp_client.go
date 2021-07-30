package infrastructure

import "github.com/cacastel/go-jsonplaceholder-middle-end/api/post/domain"

type JpClient struct {
	*JsonHttpClient
}

const (
	JSON_PLACEHOLDER_BASE_URL = "https://jsonplaceholder.typicode.com"
)

func NewJpClient() *JpClient {
	return &JpClient{
		NewJsonHttpClient(JSON_PLACEHOLDER_BASE_URL),
	}
}

func (j *JpClient) GetAllPosts() ([]*domain.Post, error) {
	posts := []*domain.Post{}
	err := j.GetJson("posts", &posts)

	return posts, err
}
