package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"goliath/logics/auth"
	"goliath/types/api"
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
		return api.NewUnauthorized(c)
	}

	token := user.(*jwt.Token)
	if token == nil {
		return api.NewUnauthorized(c)
	}

	claims := token.Claims.(*auth.Claims)
	if claims == nil {
		return api.NewUnauthorized(c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"id":   claims.UserID,
		"role": claims.Role,
	})
}
