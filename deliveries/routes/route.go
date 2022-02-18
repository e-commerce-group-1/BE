package route

import (
	"group-project1/deliveries/controllers/address"
	"group-project1/deliveries/controllers/auth"
	"group-project1/deliveries/controllers/user"
	"group-project1/deliveries/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc *user.UserController, ac *auth.AuthController, a *address.AddressController) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	e.POST("/users", uc.Insert())
	e.POST("/login", ac.Login())
	e.PUT("/users", uc.Update(), middlewares.JWTMiddleware())
	e.DELETE("/users", uc.Delete(), middlewares.JWTMiddleware())

	eAddress := e.Group("")
	eAddress.Use(middlewares.JWTMiddleware())
	eAddress.POST("/addresses", a.Insert())
	eAddress.GET("/addresses", a.Get())
	eAddress.PUT("/addresses", a.Update())
	eAddress.DELETE("/addresses", a.Delete())
}
