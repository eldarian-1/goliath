package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name"`
}

func echoUser(c echo.Context) error {
  	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/echo-user", echoUser)

	e.Logger.Fatal(e.Start(":8080"))
}
