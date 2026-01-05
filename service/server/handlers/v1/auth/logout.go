package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Logout struct{}

func (_ Logout) GetPath() string {
	return "/api/v1/auth/logout"
}

func (_ Logout) GetMethod() string {
	return http.MethodPost
}

func (_ Logout) DoHandle(c echo.Context) error {
	cookie, _ := c.Cookie("refresh")
	Service.DeleteRefresh(cookie.Value)

	ClearCookie(c, "access")
	ClearCookie(c, "refresh")

	return c.NoContent(http.StatusNoContent)
}
