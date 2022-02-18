package user

import (
	"group-project1/entities/user"

	"gorm.io/gorm"
)

// =================== Create User =======================
type CreateUserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	UserName string `json:"user_name" form:"user_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type CreateUserResponseFormat struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

func ToCreateUserResponseFormat(UserResponse user.Users) CreateUserResponseFormat {
	return CreateUserResponseFormat{
		Name:     UserResponse.Name,
		UserName: UserResponse.UserName,
		Email:    UserResponse.Email,
	}
}

// =================== Update User =======================
type UpdateUserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	UserName string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (UURF UpdateUserRequestFormat) ToUpdateUserRequestFormat(ID uint) user.Users {
	return user.Users{
		Model:    gorm.Model{ID: ID},
		Name:     UURF.Name,
		UserName: UURF.UserName,
		Email:    UURF.Email,
		Password: UURF.Password,
	}
}

type UpdateUserResponseFormat struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

func ToUpdateUserResponseFormat(UserResponse user.Users) UpdateUserResponseFormat {
	return UpdateUserResponseFormat{
		Name:     UserResponse.Name,
		UserName: UserResponse.UserName,
		Email:    UserResponse.Email,
	}
}
