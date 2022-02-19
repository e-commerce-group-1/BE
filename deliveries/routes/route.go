package route

import (
	"group-project1/deliveries/controllers/address"
	"group-project1/deliveries/controllers/admin"
	"group-project1/deliveries/controllers/auth"
	"group-project1/deliveries/controllers/product"
	"group-project1/deliveries/controllers/product_category"
	"group-project1/deliveries/controllers/user"
	"group-project1/deliveries/controllers/order"
	"group-project1/deliveries/controllers/payment_method"
	"group-project1/deliveries/controllers/transaction"

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
	r.PUT("/addresses/:id", a.Update())
	r.DELETE("/addresses/:id", a.Delete())
}

func RegisterAdminPath(e *echo.Echo, ad *admin.AdminController) {
	e.POST("/admins", ad.Insert())
	e.GET("/admins", ad.Get(), middlewares.JWTMiddleware())
}

func RegisterProductPath(e *echo.Echo, pc *product.ProductController) {
	e.POST("/products/:id", pc.Insert())
	e.GET("/products", pc.Get())
	e.GET("/products/:id", pc.GetByID())
	e.PUT("/products/:id", pc.Update())
	e.DELETE("/products/:id", pc.Delete())
}

func RegisterProductCatPath(e *echo.Echo, pc *product_category.ProdCatController) {
	e.POST("/productcats", pc.Insert())
	e.GET("/productcats", pc.Get())
	e.PUT("/productcats/:id", pc.Update())
	e.DELETE("/productcats/:id", pc.Delete())
}

func RegisterPayMethodPath(e *echo.Echo, pm *payment_method.PMController) {
	e.POST("/paymentmethods", pm.Insert())
	e.GET("/paymentmethods", pm.Get())
	e.PUT("/paymentmethods/:id", pm.Update())
	e.DELETE("/paymentmethods/:id", pm.Delete())
}

func RegisterOrderPath(e *echo.Echo, o *order.OrderController) {
	e.POST("/orders", o.Insert())
	e.GET("/orders", o.Get())
	// e.PUT("/orders/:id", pm.Update())
	// e.DELETE("/orders/:id", pm.Delete())
}

func RegisterTransactionPath(e *echo.Echo, tr *transaction.TransactionController) {
	e.POST("/transactions", tr.Insert())
	e.GET("/transactions", tr.Get())
	e.PUT("/transactions/:id", tr.Update())
	e.DELETE("/transactions/:id", tr.Delete())
}