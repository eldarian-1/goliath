package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"goliath/logics/auth"
)

var Service *auth.Service

func init() {
	Service = auth.NewService()
}

func SetCookie(c echo.Context, name, value string, maxAge int) {
	c.SetCookie(&http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   maxAge,
		SameSite: http.SameSiteLaxMode,
	})
}

func ClearCookie(c echo.Context, name string) {
	c.SetCookie(&http.Cookie{
		Name:   name,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
