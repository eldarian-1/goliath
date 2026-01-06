package auth

import (
	"net/http"
	"strconv"

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
		Name        string   `json:"name"`
		Email       string   `json:"email"`
		Password    string   `json:"password"`
		Permissions []string `json:"permissions,omitempty"`
	}

	if err := c.Bind(&req); err != nil {
		return api.NewBadRequest(c, "Bad request")
	}

	// Если name не передан, используем email как name
	if req.Name == "" {
		req.Name = req.Email
	}

	user, err := Service.Register(c.Request().Context(), req.Name, req.Email, req.Password, req.Permissions)
	if err != nil {
		return api.NewBadRequest(c, err.Error())
	}

	access, _ := auth.GenerateAccessToken(*user)
	refresh, _ := auth.GenerateRefreshToken(*user)

	Service.SaveRefresh(refresh, strconv.FormatInt(user.ID, 10))

	SetCookie(c, "access", access, 900)
	SetCookie(c, "refresh", refresh, 2592000)

	return c.NoContent(http.StatusNoContent)
}
