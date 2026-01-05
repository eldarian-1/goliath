package files

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"goliath/models/redis"
	"goliath/models/s3"
	"goliath/types/api"
)

type FilesDelete struct{}

func (_ FilesDelete) GetPath() string {
	return "/api/v1/files"
}

func (_ FilesDelete) GetMethod() string {
	return http.MethodDelete
}

func (_ FilesDelete) DoHandle(c echo.Context) error {
	fileName, ok := getFileName(c)
	if !ok {
		return api.NewBadRequest(c, "file_name is required query param")
	}

	redis.Del(c.Request().Context(), fileName)

	err := s3.Delete(c.Request().Context(), fileName)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
