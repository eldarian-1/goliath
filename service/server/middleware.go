package server

import (
	"github.com/labstack/echo/v4"
)

type Middleware interface {
	GetMiddleware() echo.MiddlewareFunc
}
