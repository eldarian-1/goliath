package cache

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"goliath/models/redis"
	"goliath/types/api"
)

type CacheGet struct{}

func (_ CacheGet) GetPath() string {
	return "/api/v1/cache"
}

func (_ CacheGet) GetMethod() string {
	return http.MethodGet
}

func (_ CacheGet) DoHandle(c echo.Context) error {
	key, ok := getKey(c)
	if !ok {
		return api.NewBadRequest(c, "key is required query param")
	}

	value, err := redis.Get(c.Request().Context(), key)
	if err != nil {
		return err
	}

	return c.Blob(http.StatusOK, "application/octet-stream", value)
}
