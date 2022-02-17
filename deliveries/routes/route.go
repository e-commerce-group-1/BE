package route

import (
	"group-project1/configs"
	"group-project1/deliveries/controllers/auth"
	"group-project1/deliveries/controllers/user"
	"group-project1/deliveries/controllers/address"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc *user.UserController, ac *auth.AuthController, a *address.AddressController) {
	e.Use(middleware.Logger())

	e.POST("/login", ac.Login())
	e.POST("/users", uc.Insert())
	e.PUT("/users", uc.Update())
	e.DELETE("/users", uc.Delete())
	e.GET("/users", uc.Get())
	
	eAddress := e.Group("")
	eAddress.Use(middleware.JWT([]byte(configs.JWT_SECRET)))
	eAddress.POST("/addresses", a.Insert())
	eAddress.GET("/addresses", a.Get())
	eAddress.PUT("/addresses", a.Update())
	eAddress.DELETE("/addresses", a.Delete())

}
