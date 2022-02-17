package auth

import "group-project1/entities/user"

type LoginRequestFormat struct {
	Email string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
}

type LoginResponseFormat struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string        `json:"message"`
	Data    user.Users `json:"data"`
	Token   string        `json:"token"`
}