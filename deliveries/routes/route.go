package route

import (
	"group-project1/deliveries/controllers/address"
	"group-project1/deliveries/controllers/admin"
	"group-project1/deliveries/controllers/auth"
	"group-project1/deliveries/controllers/order"
	"group-project1/deliveries/controllers/payment_method"
	"group-project1/deliveries/controllers/product"
	"group-project1/deliveries/controllers/transaction"
	"group-project1/deliveries/controllers/user"

	"group-project1/deliveries/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterUserPath(e *echo.Echo, uc *user.UserController) {
	e.Use(middleware.CORS())
	// e.Use(middleware.HTTPSRedirect())
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
	e.POST("/products", pc.Insert(), middlewares.JWTMiddleware())
	e.GET("/products", pc.Get())
	e.GET("/products/:id", pc.GetByID())
	e.PUT("/products/:id", pc.Update(), middlewares.JWTMiddleware())
	e.DELETE("/products/:id", pc.Delete(), middlewares.JWTMiddleware())
}

func RegisterPayMethodPath(e *echo.Echo, pm *payment_method.PMController) {
	e.POST("/paymentmethods", pm.Insert(), middlewares.JWTMiddleware())
	e.GET("/paymentmethods", pm.Get(), middlewares.JWTMiddleware())
	e.PUT("/paymentmethods/:id", pm.Update(), middlewares.JWTMiddleware())
	e.DELETE("/paymentmethods/:id", pm.Delete(), middlewares.JWTMiddleware())
}

func RegisterOrderPath(e *echo.Echo, o *order.OrderController) {
	e.POST("/orders", o.Insert(), middlewares.JWTMiddleware())
	e.GET("/orders/getuserbyid", o.GetByUserID(), middlewares.JWTMiddleware())
	e.GET("/orders/setpayed/:id", o.SetPayed(), middlewares.JWTMiddleware())
	e.GET("/orders/setcancel/:id", o.SetCancel(), middlewares.JWTMiddleware())
	e.GET("/orders", o.GetHistoryByUserID(), middlewares.JWTMiddleware())
}

func RegisterTransactionPath(e *echo.Echo, tr *transaction.TransactionController) {
	e.POST("/transactions", tr.Insert(), middlewares.JWTMiddleware())
	e.GET("/transactions", tr.GetAllTrxByUserID(), middlewares.JWTMiddleware())
	e.GET("/transactions/:id", tr.FindID(), middlewares.JWTMiddleware())
	e.DELETE("/transactions/:id", tr.DeleteByID(), middlewares.JWTMiddleware())
}
