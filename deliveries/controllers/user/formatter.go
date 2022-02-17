package user

import "group-project1/entities/user"

// =================== Create User =======================
type CreateUserRequestFormat struct {
	Name string `json:"name" form:"name"`
	User_name string `json:"user_name" form:"user_name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type CreateUserResponseFormat struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string        `json:"message"`
	Data    user.Users `json:"data"`
}

// =================== Get Users =======================
type GetUsersResponseFormat struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string        `json:"message"`
	Data   []user.Users `json:"data"`
}

// =================== Update User =======================
type UpdateUserRequestFormat struct {
	Name string `json:"name" form:"name"`
	User_name string `json:"user_name" form:"user_name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateUserResponseFormat struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string        `json:"message"`
	Data    user.Users `json:"data"`
}

// =================== Delete User =======================
type DeleteUserResponseFormat struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string        `json:"message"`
	Data   []user.Users `json:"data"`
}