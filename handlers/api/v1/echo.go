package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"goliath/types/api"
)

type Echo struct {}

func (_ Echo) GetPath() string {
	return "/api/v1/echo"
}

func (_ Echo) GetMethod() string {
	return http.MethodPost
}

func (_ Echo) DoHandle(c echo.Context) error {
	u := new(api.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
