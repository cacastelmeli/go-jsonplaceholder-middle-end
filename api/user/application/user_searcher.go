package application

import "github.com/cacastel/go-jsonplaceholder-middle-end/api/user/domain"

type UserSearcher struct {
	Repo domain.UserRepository
}

func (u *UserSearcher) SearchAll() []*domain.User {
	return u.Repo.GetUsers()
}

func (u *UserSearcher) Search(id int) *domain.User {
	return u.Repo.GetUser(id)
}
