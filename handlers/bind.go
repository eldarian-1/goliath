package handlers

import (
	"github.com/labstack/echo/v4"

	"goliath/handlers/api/v1"
)

func Bind(e *echo.Echo) {
	handlers := []Handler{
		v1.Echo{},
	}
	for _, h := range handlers {
		e.Add(h.GetMethod(), h.GetPath(), h.DoHandle)
	}
}
