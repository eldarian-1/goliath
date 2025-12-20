package users

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"goliath/types/api"
	"goliath/repositories"
)

const (
	defaultLimit = 10
	defaultWithDeleted = false
)

type UsersGet struct {}

func (_ UsersGet) GetPath() string {
	return "/api/v1/users"
}

func (_ UsersGet) GetMethod() string {
	return http.MethodGet
}

func (_ UsersGet) DoHandle(c echo.Context) error {
	postgresUsers, err := repositories.GetUsers(
		getLimit(c),
		getCursorById(c),
		getWithDeleted(c),
	)
	if err != nil {
		return err
	}

	var response api.GetUsersResponse
	response.Users = make([]api.User, 0, len(postgresUsers))
	for _, user := range postgresUsers {
		response.Users = append(response.Users, api.User{
			Id: user.Id,
			Name: user.Name,
			CreatedAt: &user.CreatedAt,
			UpdatedAt: &user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		})
	}

	return c.JSON(http.StatusOK, response)
}

func getLimit(c echo.Context) int64 {
	limitStr := c.QueryParam("limit")
	limit, err := strconv.ParseInt(limitStr, 10, 64)

	if err != nil {
		return defaultLimit
	}

	return limit
}

func getCursorById(c echo.Context) *int64 {
	cursorStr := c.QueryParam("cursor")
	cursor, err := strconv.ParseInt(cursorStr, 10, 64)

	if err != nil {
		return nil
	}

	return &cursor
}

func getWithDeleted(c echo.Context) bool {
	withDeletedStr := c.QueryParam("with_deleted")
	withDeleted, err := strconv.ParseBool(withDeletedStr)

	if err != nil {
		return defaultWithDeleted
	}

	return withDeleted
}
