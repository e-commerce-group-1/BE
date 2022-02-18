package user

import (
	"group-project1/deliveries/controllers/common"
	"group-project1/deliveries/middlewares"
	"group-project1/entities/user"
	userRepo "group-project1/repository/user"
	"net/http"

	"github.com/labstack/echo/v4"
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
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan user baru", ToCreateUserResponseFormat(res)))
	}
}

func (uc *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)
		var UpdatedUser = UpdateUserRequestFormat{}

		if err := c.Bind(&UpdatedUser); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := uc.repo.Update(UpdatedUser.ToUpdateUserRequestFormat(uint(UserID)))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses update user", ToUpdateUserResponseFormat(res)))
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
