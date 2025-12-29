package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewBadRequest(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, Error{
		Code:	"bad_request",
		Message: message,
	})
}
