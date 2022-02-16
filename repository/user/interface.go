package user

import u "group-project1/enitities/user"

type User interface {
	Get() ([]u.Users, error)
	Insert(newUser u.Users) (u.Users, error)
	Update(userId int, newUser u.Users) (u.Users, error)
	Delete(userId int) error
}
