package auth

import (
	"group-project1/configs"
	"group-project1/entities/user"
	"group-project1/repository/auth"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repo auth.Auth
}

func New(repo auth.Auth) *AuthController {
	return &AuthController{
		repo: repo,
	}
}

func (a AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginFormat := LoginRequestFormat{}

		if err := c.Bind(&loginFormat); err != nil {
			return c.JSON(http.StatusBadRequest, "Ada yang salah dengan input")
		}

		checkedUser, err := a.repo.Login(loginFormat.Email, loginFormat.Password)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Email atau password tidak valid")
		}

		token, err := GenerateToken(checkedUser)

		if err != nil {
			return c.JSON(http.StatusNotAcceptable, "Terjadi error ketika melakukan generate token")
		}

		return c.JSON(http.StatusOK, LoginResponseFormat{
			Code: 200,
			Success: true,
			Message: "Success Login",
			Data: checkedUser,
			Token: token,
		})
	}
}

func GenerateToken(user user.Users) (string, error) {
	datas := jwt.MapClaims{}
	datas["id"] = user.ID
	datas["exp"] = time.Now().Add(time.Hour * 1).Unix() //1jam
	datas["authorized"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, datas)
	return token.SignedString([]byte(configs.JWT_SECRET))
}

func ExtractTokenUserId(e echo.Context) float64 {

	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		datas := user.Claims.(jwt.MapClaims)
		userId := datas["id"].(float64)
		return userId
	}

	return 0
}
