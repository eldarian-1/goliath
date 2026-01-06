package middlewares

import (
	"goliath/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var origin string

type CORS struct{}

func (_ CORS) GetMiddleware() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{origin},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			"Range",
		},
		ExposeHeaders: []string{
			"Content-Range",
			"Accept-Ranges",
			"Content-Length",
		},
		AllowCredentials: true,
	})
}

func init() {
	origin = utils.GetEnv("CORS_ALLOW_ORIGIN", "http://localhost:5173")
}
