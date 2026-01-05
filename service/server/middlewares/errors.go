package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"goliath/types/api"
)

type Errors struct{}

func (_ Errors) GetMiddleware() echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, api.Error{
					Code:    "internal_server_error",
					Message: err.Error(),
				})
			}

			return nil
		}
	})
}
