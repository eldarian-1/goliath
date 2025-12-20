package users

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"goliath/repositories"
	"goliath/types/api"
	"goliath/types/postgres"
)

type UsersPost struct {}

func (_ UsersPost) GetPath() string {
	return "/api/v1/users"
}

func (_ UsersPost) GetMethod() string {
	return http.MethodPost
}

func (_ UsersPost) DoHandle(c echo.Context) error {
	u := new(api.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	err := repositories.UpsertUser(postgres.User{
		Id: u.Id,
		Name: u.Name,
		DeletedAt: nil,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, u)
}
