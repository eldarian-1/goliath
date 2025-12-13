package handlers

import (
	"github.com/labstack/echo/v4"
)

type Handler interface {
	GetPath() string
	GetMethod() string
	DoHandle(c echo.Context) error
}
