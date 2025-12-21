package users

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"goliath/repositories"
	"goliath/types/api"
)

const (
	defaultLimit       = 10
	defaultWithDeleted = false
)

type UsersGet struct{}

func (_ UsersGet) GetPath() string {
	return "/api/v1/users"
}

func (_ UsersGet) GetMethod() string {
	return http.MethodGet
}

func (_ UsersGet) DoHandle(c echo.Context) error {
	limit := getLimit(c)
	postgresUsers, err := repositories.GetUsers(
		limit+1,
		getCursorById(c),
		getWithDeleted(c),
	)
	if err != nil {
		return err
	}

	var response api.GetUsersResponse
	response.Users = make([]api.User, 0, len(postgresUsers))
	for i, user := range postgresUsers {
		if i >= int(limit) {
			break
		}
		var deletedAt *time.Time
		if user.DeletedAt.Valid {
			deletedAt = &user.DeletedAt.Time
		}
		response.Users = append(response.Users, api.User{
			Id:        &user.Id.Int64,
			Name:      user.Name,
			CreatedAt: &user.CreatedAt,
			UpdatedAt: &user.UpdatedAt,
			DeletedAt: deletedAt,
		})
	}

	if int64(len(postgresUsers)) > limit {
		response.CursorById = &postgresUsers[limit].Id.Int64
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
