package application

import "github.com/cacastel/go-jsonplaceholder-middle-end/api/user/domain"

type UserSaver struct {
	Repo domain.UserRepository
}

func (u *UserSaver) Save(user *domain.User) {
	// proceso de guardado
	u.Repo.SaveUser(user)
}
