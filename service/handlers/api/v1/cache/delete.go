package cache

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"goliath/models/redis"
	"goliath/types/api"
)

type CacheDelete struct{}

func (_ CacheDelete) GetPath() string {
	return "/api/v1/cache"
}

func (_ CacheDelete) GetMethod() string {
	return http.MethodDelete
}

func (_ CacheDelete) DoHandle(c echo.Context) error {
	key, ok := getKey(c)
	if !ok {
		return api.NewBadRequest(c, "key is required query param")
	}

	redis.Del(c.Request().Context(), key)

	return c.NoContent(http.StatusNoContent)
}
