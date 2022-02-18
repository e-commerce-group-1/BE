package admin

import (
	"group-project1/deliveries/controllers/common"
	"group-project1/deliveries/middlewares"
	UserEntity "group-project1/entities/user"
	"group-project1/repository/admin"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	repo admin.Admin
}

func New(repository admin.Admin) *AdminController {
	return &AdminController{
		repo: repository,
	}
}

func (ac *AdminController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		NewUser := AdminCreateRequestFormat{}

		if err := c.Bind(&NewUser); err != nil || NewUser.Name == "" || NewUser.Email == "" || NewUser.Password == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		newUser := UserEntity.Users{
			Name:     NewUser.Name,
			UserName: NewUser.UserName,
			Email:    NewUser.Email,
			Password: NewUser.Password,
		}
		res, err := ac.repo.Insert(newUser)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan admin baru", ToAdminCreateResponseFormat(res)))
	}
}

func (uc *AdminController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized())
		}

		res, err := uc.repo.Get()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua user", ToAdminGetResponseFormat(res)))
	}
}
