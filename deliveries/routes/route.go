package route

import (
	"group-project1/deliveries/controllers/address"
	"group-project1/deliveries/controllers/admin"
	"group-project1/deliveries/controllers/auth"
	"group-project1/deliveries/controllers/product"
	"group-project1/deliveries/controllers/user"
	"group-project1/deliveries/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterUserPath(e *echo.Echo, uc *user.UserController) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	e.POST("/users", uc.Insert())
	e.PUT("/users", uc.Update(), middlewares.JWTMiddleware())
	e.DELETE("/users", uc.Delete(), middlewares.JWTMiddleware())
}

func RegisterAuthPath(e *echo.Echo, ac *auth.AuthController) {
	e.POST("/login", ac.Login())
}

func RegisterAddressPath(e *echo.Echo, a *address.AddressController) {
	r := e.Group("jwt/")
	r.Use(middlewares.JWTMiddleware())
	r.POST("/addresses", a.Insert())
	r.GET("/addresses", a.Get())
	r.PUT("/addresses", a.Update())
	r.DELETE("/addresses", a.Delete())
}

func RegisterAdminPath(e *echo.Echo, ad *admin.AdminController) {
	e.POST("/admins", ad.Insert())
	e.GET("/admins", ad.Get(), middlewares.JWTMiddleware())
}

func RegisterProductPath(e *echo.Echo, pc *product.ProductController) {
	e.POST("/products", pc.Insert())
	e.PUT("/products", pc.Update())
	e.DELETE("/products", pc.Delete())
}
