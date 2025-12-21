package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"goliath/handlers/api/v1"
	"goliath/handlers/api/v1/users"
	"goliath/types/api"
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

func MyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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
}

func Define() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(MyMiddleware)

	for _, h := range handlers {
		e.Add(h.GetMethod(), h.GetPath(), h.DoHandle)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
