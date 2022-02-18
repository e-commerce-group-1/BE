package admin

import "group-project1/entities/user"

type Admin interface {
	Insert(NewUser user.Users) (user.Users, error)
}
