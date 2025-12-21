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

	ok, err := repositories.UpsertUser(postgres.User{
		Id:        id,
		Name:      u.Name,
		DeletedAt: sql.NullTime{Valid: false},
	})

	if err != nil {
		return err
	}

	if !ok {
		return c.JSON(http.StatusNotFound, api.Error{
			Code:    "not_found",
			Message: "User not found",
		})
	}

	return c.JSON(http.StatusNoContent, u)
}
