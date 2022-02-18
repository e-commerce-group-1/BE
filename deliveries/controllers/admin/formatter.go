package admin

import "group-project1/entities/user"

type AdminCreateRequestFormat struct {
	Name     string `json:"name" form:"name"`
	UserName string `json:"user_name" form:"user_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type AdminCreateResponseFormat struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

func ToAdminCreateResponseFormat(AdminResponse user.Users) AdminCreateResponseFormat {
	return AdminCreateResponseFormat{
		Name:     AdminResponse.Name,
		UserName: AdminResponse.UserName,
		Email:    AdminResponse.Email,
	}
}

type AdminGetResponseFormat struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

func ToAdminGetResponseFormat(AdminResponses []user.Users) []AdminGetResponseFormat {
	AdminGetResponses := make([]AdminGetResponseFormat, len(AdminResponses))
	for i := 0; i < len(AdminResponses); i++ {
		AdminGetResponses[i].Name = AdminResponses[i].Name
		AdminGetResponses[i].UserName = AdminResponses[i].UserName
		AdminGetResponses[i].Email = AdminResponses[i].Email
	}
	return AdminGetResponses
}
