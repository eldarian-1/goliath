package files

import (
	"github.com/labstack/echo/v4"
)

func getFileName(c echo.Context) (string, bool) {
	fileName := c.QueryParam("name")

	return fileName, fileName != ""
}
