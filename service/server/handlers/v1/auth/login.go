package auth

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"goliath/logics/auth"
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

	user, err := Service.Login(c.Request().Context(), req.Email, req.Password)
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
