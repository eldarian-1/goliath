package handlers

import (
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

type Metrics struct{}

func (_ Metrics) GetPath() string {
	return "/metrics"
}

func (_ Metrics) GetMethod() string {
	return http.MethodGet
}

func (_ Metrics) DoHandle(c echo.Context) error {
	return echoprometheus.NewHandler()(c)
}
