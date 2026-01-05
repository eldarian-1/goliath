package users

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"goliath/repositories"
	"goliath/types/api"
	"goliath/types/postgres"
)

type UsersPost struct{}

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

	id := sql.NullInt64{Valid: u.Id != nil}
	if u.Id != nil {
		id.Int64 = *u.Id
	}

	_, err := repositories.UpsertUser(
		c.Request().Context(),
		postgres.User{
			Id:        id,
			Name:      u.Name,
			DeletedAt: sql.NullTime{Valid: false},
		},
	)

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
