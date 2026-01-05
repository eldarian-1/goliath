package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"goliath/services/auth"
	"goliath/types/api"
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
		return api.NewUnauthorized(c)
	}

	_, ok := Service.ValidateRefresh(cookie.Value)
	if !ok {
		return api.NewUnauthorized(c)
	}

	t, err := jwt.ParseWithClaims(cookie.Value,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return auth.RefreshSecret, nil
		})

	if err != nil || !t.Valid {
		return api.NewUnauthorized(c)
	}

	userID := t.Claims.(*jwt.RegisteredClaims).Subject
	user := Service.Users[userID]

	access, _ := auth.GenerateAccessToken(user)
	SetCookie(c, "access", access, 900)

	return c.NoContent(http.StatusNoContent)
}
