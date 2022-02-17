package user

import "group-project1/entities/user"

// =================== Create User =======================
type CreateUserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	UserName string `json:"user_name" form:"user_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type CreateUserResponseFormat struct {
	Code    int        `json:"code"`
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    user.Users `json:"data"`
}

// =================== Update User =======================
type UpdateUserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	UserName string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
