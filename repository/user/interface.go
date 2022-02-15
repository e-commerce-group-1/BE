package user

import u "group-project1/enitities/user"

type User interface {
	Get() ([]u.User, error)
	Insert(newUser u.User) (u.User, error)
	Update(userId int, newUser u.User) (u.User, error)
	Delete(userId int) error
}