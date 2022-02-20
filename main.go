package main

import (
	"group-project1/configs"
	"group-project1/deliveries/controllers/address"
	"group-project1/deliveries/controllers/admin"
	"group-project1/deliveries/controllers/auth"
	"group-project1/deliveries/controllers/order"
	"group-project1/deliveries/controllers/payment_method"
	"group-project1/deliveries/controllers/product"
	"group-project1/deliveries/controllers/user"
	route "group-project1/deliveries/routes"
	"group-project1/utils"

	_addressRepo "group-project1/repository/address"
	_adminRepo "group-project1/repository/admin"
	_authRepo "group-project1/repository/auth"
	_orderRepo "group-project1/repository/order"
	_payRepo "group-project1/repository/payment_method"
	_productRepo "group-project1/repository/product"
	_userRepo "group-project1/repository/user"

	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	userRepo := _userRepo.New(db)
	userController := user.New(userRepo)
	authRepo := _authRepo.New(db)
	authController := auth.New(authRepo)
	addressRepo := _addressRepo.New(db)
	addressController := address.New(addressRepo)
	adminRepo := _adminRepo.New(db)
	adminController := admin.New(adminRepo)
	productRepo := _productRepo.New(db)
	productController := product.New(productRepo)
	orderRepo := _orderRepo.New(db)
	orderController := order.New(orderRepo)
	payMethodRepo := _payRepo.New(db)
	payMethodController := payment_method.New(payMethodRepo)
	// transactionRepo := _transactionRepo.New(db)
	// transactionController := transaction.New(transactionRepo)

	e := echo.New()

	route.RegisterUserPath(e, userController)
	route.RegisterAuthPath(e, authController)
	route.RegisterAddressPath(e, addressController)
	route.RegisterAdminPath(e, adminController)
	route.RegisterProductPath(e, productController)
	route.RegisterOrderPath(e, orderController)
	route.RegisterPayMethodPath(e, payMethodController)
	// route.RegisterTransactionPath(e, transactionController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
