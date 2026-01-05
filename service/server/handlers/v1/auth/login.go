package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"goliath/server/handlers/v1/auth/gpt"
	"goliath/types/api"
)

type Login struct{}

func (_ Login) GetPath() string {
	return "/api/v1/auth/login"
}

func (_ Login) GetMethod() string {
	return http.MethodPost
}

func (_ Login) DoHandle(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		return api.NewBadRequest(c, "Bad request")
	}

	user, err := Service.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}

	access, _ := gpt.GenerateAccessToken(*user)
	refresh, _ := gpt.GenerateRefreshToken(*user)

	Service.SaveRefresh(refresh, user.ID)

	SetCookie(c, "access", access, 900)
	SetCookie(c, "refresh", refresh, 2592000)

	return c.NoContent(http.StatusNoContent)
}
