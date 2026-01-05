package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"goliath/server/handlers/v1/auth/gpt"
)

type Refresh struct{}

func (_ Refresh) GetPath() string {
	return "/api/v1/auth/refresh"
}

func (_ Refresh) GetMethod() string {
	return http.MethodPost
}

func (_ Refresh) DoHandle(c echo.Context) error {
	cookie, err := c.Cookie("refresh")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}

	_, ok := Service.ValidateRefresh(cookie.Value)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "user not found"})
	}

	t, err := jwt.ParseWithClaims(cookie.Value,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return gpt.RefreshSecret, nil
		})

	if err != nil || !t.Valid {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "token is not valid"})
	}

	userID := t.Claims.(*jwt.RegisteredClaims).Subject
	user := Service.Users[userID]

	access, _ := gpt.GenerateAccessToken(user)
	SetCookie(c, "access", access, 900)

	return c.NoContent(http.StatusNoContent)
}
