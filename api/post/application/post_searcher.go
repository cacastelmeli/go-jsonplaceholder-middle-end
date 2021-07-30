package application

import "github.com/cacastel/go-jsonplaceholder-middle-end/api/post/domain"

type PostSearcher struct {
	Repo domain.PostRepository
}

func (p *PostSearcher) SearchAll() ([]*domain.Post, error) {
	return p.Repo.FetchAllPosts()
}

func (p *PostSearcher) Search(limit uint) ([]*domain.Post, error) {
	allPosts, err := p.SearchAll()

	if err != nil {
		return nil, err
	}

	return allPosts[:limit], nil
}
