package main

import (
	"group-project1/configs"
	"group-project1/deliveries/controllers/address"
	"group-project1/deliveries/controllers/admin"
	"group-project1/deliveries/controllers/auth"
	"group-project1/deliveries/controllers/user"
	route "group-project1/deliveries/routes"
	"group-project1/utils"

	_addressRepo "group-project1/repository/address"
	_adminRepo "group-project1/repository/admin"
	_authRepo "group-project1/repository/auth"
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

	e := echo.New()

	route.RegisterUserPath(e, userController)
	route.RegisterAuthPath(e, authController)
	route.RegisterAddressPath(e, addressController)
	route.RegisterAdminPath(e, adminController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
