package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewUnauthorized(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, Error{
		Code:    "unauthorized",
		Message: "Unauthorized user",
	})
}

func NewBadRequest(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, Error{
		Code:    "bad_request",
		Message: message,
	})
}
