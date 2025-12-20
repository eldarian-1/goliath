package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"goliath/handlers/api/v1"
	"goliath/handlers/api/v1/users"
)

var handlers []Handler

func init() {
	handlers = []Handler{
		users.UsersGet{},
		users.UsersPost{},
		users.UsersDelete{},
		v1.Log{},
	}
}

func Define() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	for _, h := range handlers {
		e.Add(h.GetMethod(), h.GetPath(), h.DoHandle)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
