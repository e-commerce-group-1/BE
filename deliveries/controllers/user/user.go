package user

import (
	"net/http"
	"group-project1/entities/user"
	userRepo "group-project1/repository/user"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserController struct {
	Repo userRepo.User
}

func New(user userRepo.User) *UserController {
	return &UserController{Repo: user}
}

func (uc UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uc.Repo.Get()

		if err != nil {
			log.Info("Got error here")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Something wrong",
			})
		}
		return c.JSON(http.StatusOK, GetUsersResponseFormat{
			Code: 200,
			Success: true,
			Message: "Success Get All User",
			Data: users,
		})
	}
}

func (uc UserController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestFormat := CreateUserRequestFormat{}

		if err := c.Bind(&requestFormat); err != nil {
			return c.JSON(http.StatusBadRequest, "Ada yang salah dengan input")
		}

		newUser := user.Users{
			Name: requestFormat.Name,
			User_name: requestFormat.User_name,
			Email: requestFormat.Email,
			Password: requestFormat.Password,
		}

		res, err := uc.Repo.Insert(newUser)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Gagal Memasukkan Data User")
		}

		return c.JSON(http.StatusOK, CreateUserResponseFormat{
			Code: 200,
			Success: true,
			Message: "Success Create User",
			Data: res,
		})
	}
}

func (uc *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newUser = UpdateUserRequestFormat{}
		userId, _ := strconv.Atoi(c.Param("id"))

		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, "Ada yang salah dengan input")
		}

		res, err := uc.Repo.Update(userId, user.Users{
			Name: newUser.Name,
			User_name: newUser.User_name,
			Email: newUser.Email,
			Password: newUser.Password,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Gagal Memperbaharui Data User")
		}

		return c.JSON(http.StatusOK, UpdateUserResponseFormat{
			Code: 200,
			Success: true,
			Message: "Success Update User",
			Data: res,
		})
	}
}

func (uc *UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, _ := strconv.Atoi(c.Param("id"))

		err := uc.Repo.Delete(userId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Gagal Menghapus User")
		}

		return c.JSON(http.StatusOK, UpdateUserResponseFormat{
			Code: 200,
			Success: true,
			Message: "Success Update User",
		})
	}
}
