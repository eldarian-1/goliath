package files

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"

	"goliath/models/s3"
	"goliath/types/api"
)

type FilesPut struct{}

func (_ FilesPut) GetPath() string {
	return "/api/v1/files"
}

func (_ FilesPut) GetMethod() string {
	return http.MethodPut
}

func (_ FilesPut) DoHandle(c echo.Context) error {
	file, err := getFile(c)
	if err != nil {
		return err
	}

	err = s3.Put(c.Request().Context(), file)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, nil)
}

func getFile(c echo.Context) (*s3.File, error) {
	contentType, ok := getHeader(c, "Content-Type")
	if !ok {
		return nil, api.NewBadRequest(c, "Content-Type is required header")
	}

	contentDisposition, ok := getHeader(c, "Content-Disposition")
	if !ok {
		return nil, api.NewBadRequest(c, "Content-Disposition is required header")
	}

	regexp, err := regexp.Compile("(filename\\s*=\\s*\"?([^\";]+)\"?)")
	if err != nil {
		return nil, err
	}

	fileName := regexp.FindStringSubmatch(contentDisposition)[2]

	file := s3.File{
		Name: fileName,
		ContentType: contentType,
		ContentDisposition: contentDisposition,
		Reader: c.Request().Body,
	}

	return &file, nil
}

func getHeader(c echo.Context, key string) (string, bool) {
	header := c.Request().Header.Get(key)
	return header, header != ""
}
