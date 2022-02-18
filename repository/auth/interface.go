package auth

import u "group-project1/entities/user"

type Auth interface {
	Login(email, password string) (u.Users, error)
}
