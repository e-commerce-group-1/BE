package user

import (
	"group-project1/deliveries/controllers/common"
	"group-project1/deliveries/middlewares"
	"group-project1/entities/user"
	userRepo "group-project1/repository/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	repo userRepo.User
}

func New(repository userRepo.User) *UserController {
	return &UserController{
		repo: repository,
	}
}

func (uc *UserController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		NewUser := CreateUserRequestFormat{}

		if err := c.Bind(&NewUser); err != nil || NewUser.Name == "" || NewUser.Email == "" || NewUser.Password == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		newUser := user.Users{
			Name:     NewUser.Name,
			UserName: NewUser.UserName,
			Email:    NewUser.Email,
			Password: NewUser.Password,
		}
		res, err := uc.repo.Insert(newUser)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan user baru", res))
	}
}

func (uc *UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		// UserID := middlewares.ExtractTokenUserId(c)
		res, err := uc.repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua user", res))
	}
}

func (uc *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)
		var UpdatedUser = UpdateUserRequestFormat{}

		if err := c.Bind(&UpdatedUser); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := uc.repo.Update(user.Users{
			Model:    gorm.Model{ID: uint(UserID)},
			Name:     UpdatedUser.Name,
			UserName: UpdatedUser.UserName,
			Email:    UpdatedUser.Email,
			Password: UpdatedUser.Password,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses update user", res))
	}
}

func (uc *UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)
		err := uc.repo.Delete(UserID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menghapus user", err))
	}
}
