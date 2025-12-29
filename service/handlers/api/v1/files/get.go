package files

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"goliath/models/s3"
	"goliath/types/api"
)

type FilesGet struct{}

func (_ FilesGet) GetPath() string {
	return "/api/v1/files"
}

func (_ FilesGet) GetMethod() string {
	return http.MethodGet
}

func (_ FilesGet) DoHandle(c echo.Context) error {
	fileName, ok := getFileName(c)
	if !ok {
		return api.NewBadRequest(c, "file_name is required query param")
	}

	file, err := s3.Get(c.Request().Context(), fileName)
	if err != nil {
		return err
	}

	return c.Stream(http.StatusOK, file.ContentType, file.Reader)
}
