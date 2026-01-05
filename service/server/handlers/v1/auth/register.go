package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"goliath/server/handlers/v1/auth/gpt"
)

type Register struct{}

func (_ Register) GetPath() string {
	return "/api/v1/auth/register"
}

func (_ Register) GetMethod() string {
	return http.MethodPost
}

func (_ Register) DoHandle(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := Service.Register(req.Email, req.Password)
	if err != nil {
		return err
	}

	access, _ := gpt.GenerateAccessToken(*user)
	refresh, _ := gpt.GenerateRefreshToken(*user)

	Service.SaveRefresh(refresh, user.ID)

	SetCookie(c, "access", access, 900)
	SetCookie(c, "refresh", refresh, 2592000)

	return c.NoContent(http.StatusNoContent)
}
