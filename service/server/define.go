package server

import (
	"fmt"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	h "goliath/server/handlers"
	"goliath/server/handlers/v1"
	"goliath/server/handlers/v1/auth"
	"goliath/server/handlers/v1/cache"
	"goliath/server/handlers/v1/files"
	"goliath/server/handlers/v1/users"
	"goliath/server/handlers/v1/videos"
	mw "goliath/server/middlewares"
	"goliath/utils"
)

var handlers []Handler
var middlewares []Middleware
var port string

func init() {
	middlewares = []Middleware{
		mw.CORS{},
		mw.Errors{},
		mw.JWT{},
	}
	handlers = []Handler{
		auth.Login{},
		auth.Logout{},
		auth.Me{},
		auth.Refresh{},
		auth.Register{},
		h.Metrics{},
		cache.CacheGet{},
		cache.CachePost{},
		cache.CacheDelete{},
		files.FilesGet{},
		files.FilesPut{},
		files.FilesDelete{},
		users.UsersGet{},
		users.UsersPost{},
		users.UsersDelete{},
		videos.Upload{},
		videos.List{},
		videos.Get{},
		v1.Log{},
	}
	port = fmt.Sprintf(":%s", utils.GetEnv("GOLIATH_PORT", "8080"))
}

func Define() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(echoprometheus.NewMiddleware("goliath"))

	for _, m := range middlewares {
		e.Use(m.GetMiddleware())
	}

	for _, h := range handlers {
		e.Add(h.GetMethod(), h.GetPath(), h.DoHandle)
	}

	e.Logger.Fatal(e.Start(port))
}
