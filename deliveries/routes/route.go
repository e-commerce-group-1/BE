package route

import (
	"group-project1/deliveries/controllers/auth"
	"group-project1/deliveries/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc *user.UserController, ac *auth.AuthController) {
	e.Use(middleware.Logger())

	e.GET("/users", uc.Get())

	e.POST("/login", ac.Login())

	e.POST("/users", uc.Insert())

}
