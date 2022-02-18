package middlewares

import (
	"errors"
	"fmt"
	"group-project1/configs"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateToken(ID uint, isAdmin bool) (string, error) {
	if ID < 1 {
		return "", errors.New("user_id is not valid")
	}
	data := jwt.MapClaims{}
	data["id"] = ID
	data["isAdmin"] = isAdmin
	data["expired"] = time.Now().Add(time.Hour * 1).Unix()
	data["authorized"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString([]byte(configs.JWT_SECRET))
}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		data := user.Claims.(jwt.MapClaims)
		fmt.Println(data["id"])
		id := int(data["id"].(float64))
		return id
	}
	return 0
}

func ExtractTokenIsAdmin(e echo.Context) bool {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		data := user.Claims.(jwt.MapClaims)
		fmt.Println(data["isAdmin"])
		isAdmin := data["isAdmin"].(bool)
		return isAdmin
	}
	return false
}
