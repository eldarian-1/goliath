package cache

import (
	"github.com/labstack/echo/v4"
)

func getKey(c echo.Context) (string, bool) {
	key := c.QueryParam("key")

	return key, key != ""
}
