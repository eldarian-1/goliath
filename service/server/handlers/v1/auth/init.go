package auth

import (
	"goliath/server/handlers/v1/auth/gpt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var Service *gpt.Service

func init() {
	Service = gpt.NewService()
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
