package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"goliath/logics/auth"
	"goliath/types/api"
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
		return api.NewBadRequest(c, "Bad request")
	}

	user, err := Service.Register(req.Email, req.Password)
	if err != nil {
		return api.NewBadRequest(c, err.Error())
	}

	access, _ := auth.GenerateAccessToken(*user)
	refresh, _ := auth.GenerateRefreshToken(*user)

	Service.SaveRefresh(refresh, user.ID)

	SetCookie(c, "access", access, 900)
	SetCookie(c, "refresh", refresh, 2592000)

	return c.NoContent(http.StatusNoContent)
}
