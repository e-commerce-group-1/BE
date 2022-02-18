package user

import u "group-project1/entities/user"

type User interface {
	Insert(newUser u.Users) (u.Users, error)
	Update(userUpdate u.Users) (u.Users, error)
	Delete(ID int) error
}
