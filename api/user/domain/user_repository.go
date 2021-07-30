package domain

type UserRepository interface {
	SaveUser(user *User)
	GetUser(id int) *User
	GetUsers() []*User
}
