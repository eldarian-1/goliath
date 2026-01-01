package cache

import (
	"io"	
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"goliath/models/redis"
	"goliath/types/api"
)

type CachePost struct{}

func (_ CachePost) GetPath() string {
	return "/api/v1/cache"
}

func (_ CachePost) GetMethod() string {
	return http.MethodPost
}

func (_ CachePost) DoHandle(c echo.Context) error {
	key, ok := getKey(c)
	if !ok {
		return api.NewBadRequest(c, "key is required query param")
	}

	reader := c.Request().Body
	defer reader.Close()

	value, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	err = redis.Set(c.Request().Context(), key, value, 5*time.Minute)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
