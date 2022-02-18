package admin

import "group-project1/entities/user"

type CreateAdminRequestFormat struct {
	Name     string `json:"name" form:"name"`
	UserName string `json:"user_name" form:"user_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type CreateAdminResponseFormat struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

func ToCreateAdminResponseFormat(AdminResponse user.Users) CreateAdminResponseFormat {
	return CreateAdminResponseFormat{
		Name:     AdminResponse.Name,
		UserName: AdminResponse.UserName,
		Email:    AdminResponse.Email,
	}
}
