package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"goliath/server/handlers/v1/auth/gpt"
)

type Me struct{}

func (_ Me) GetPath() string {
	return "/api/v1/auth/me"
}

func (_ Me) GetMethod() string {
	return http.MethodGet
}

func (_ Me) DoHandle(c echo.Context) error {
	user := c.Get("user")
	if user == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "user is nil",
		})
	}

	token := user.(*jwt.Token)
	if token == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "token is nil",
		})
	}

	claims := token.Claims.(*gpt.Claims)
	if claims == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "claims is nil",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"id":   claims.UserID,
		"role": claims.Role,
	})
}
