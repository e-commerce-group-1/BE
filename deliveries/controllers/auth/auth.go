package auth

import (
	"group-project1/deliveries/controllers/common"
	"group-project1/deliveries/middlewares"
	"group-project1/repository/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repo auth.Auth
}

func New(repository auth.Auth) *AuthController {
	return &AuthController{
		repo: repository,
	}
}

func (a *AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginFormat := LoginRequestFormat{}
		if err := c.Bind(&loginFormat); err != nil || loginFormat.Email == "" || loginFormat.Password == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		checkedUser, err := a.repo.Login(loginFormat.Email, loginFormat.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}

		isAdmin := checkedUser.IsAdmin
		tokenID, err := middlewares.GenerateToken(checkedUser.ID, isAdmin)
		if err != nil {
			return c.JSON(http.StatusNotAcceptable, common.NotAcceptable())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "berhasil masuk, mendapatkan token baru", tokenID))
	}
}
