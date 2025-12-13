package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"goliath/handlers"
)

func main() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	handlers.Bind(e)

	e.Logger.Fatal(e.Start(":8080"))
}
