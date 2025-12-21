package users

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"goliath/repositories"
	"goliath/types/api"
)

type UsersDelete struct{}

func (_ UsersDelete) GetPath() string {
	return "/api/v1/users"
}

func (_ UsersDelete) GetMethod() string {
	return http.MethodDelete
}

func (_ UsersDelete) DoHandle(c echo.Context) error {
	id, err := getUserId(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.Error{
			Code:    "bad_request",
			Message: "Invalid id",
		})
	}

	users, err := repositories.GetUsers(1, &id, false)
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, api.Error{
			Code:    "not_found",
			Message: "User not found",
		})
	}

	user := users[0]
	user.DeletedAt.Valid = true
	user.DeletedAt.Time = time.Now()

	_, err = repositories.UpsertUser(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, nil)
}

func getUserId(c echo.Context) (int64, error) {
	idStr := c.QueryParam("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		return 0, err
	}

	return id, nil
}
