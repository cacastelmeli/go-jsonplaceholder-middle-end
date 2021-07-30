package infrastructure

import (
	"github.com/cacastel/go-jsonplaceholder-middle-end/api/user/domain"
)

type FakeUserRepository struct {
	users []*domain.User
}

func (f *FakeUserRepository) SaveUser(user *domain.User) {
	f.users = append(f.users, user)
}

func (f *FakeUserRepository) GetUser(id int) *domain.User {
	for _, user := range f.users {
		if user.Id == id {
			return user
		}
	}

	return nil
}

func (f *FakeUserRepository) GetUsers() []*domain.User {
	return f.users
}
