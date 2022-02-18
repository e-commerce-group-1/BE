package main

import (
	"group-project1/configs"
	"group-project1/deliveries/controllers/auth"
	"group-project1/deliveries/controllers/user"
	"group-project1/deliveries/controllers/address"
	"group-project1/deliveries/routes"
	"group-project1/utils"

	_authRepo "group-project1/repository/auth"
	_userRepo "group-project1/repository/user"
	_addressRepo "group-project1/repository/address"

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

	e := echo.New()

	route.RegisterPath(e, userController, authController, addressController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}