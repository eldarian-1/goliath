package files

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"goliath/caches"
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

	var file *s3.File
	var err error

	if cached(c) {
		file, err = caches.Files{}.Get(c.Request().Context(), fileName)
	} else {
		file, err = s3.Get(c.Request().Context(), fileName)
	}

	if err != nil {
		return err
	}

	return c.Stream(http.StatusOK, file.ContentType, file.Reader)
}

func cached(c echo.Context) bool {
	return c.QueryParam("cache") == "true"
}
