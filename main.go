package main

import (
	"group-project1/configs"
	"group-project1/deliveries/controllers/address"
	"group-project1/deliveries/controllers/admin"
	"group-project1/deliveries/controllers/auth"
	"group-project1/deliveries/controllers/product"
	"group-project1/deliveries/controllers/product_category"
	"group-project1/deliveries/controllers/user"
	route "group-project1/deliveries/routes"
	"group-project1/utils"

	_addressRepo "group-project1/repository/address"
	_adminRepo "group-project1/repository/admin"
	_authRepo "group-project1/repository/auth"
	_productRepo "group-project1/repository/product"
	_productCatRepo "group-project1/repository/product_category"
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
	productCatRepo := _productCatRepo.New(db)
	productCatController := product_category.New(productCatRepo)
	e := echo.New()

	route.RegisterUserPath(e, userController)
	route.RegisterAuthPath(e, authController)
	route.RegisterAddressPath(e, addressController)
	route.RegisterAdminPath(e, adminController)
	route.RegisterProductPath(e, productController)
	route.RegisterProductCatPath(e, productCatController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
