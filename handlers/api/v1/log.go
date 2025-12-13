package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"goliath/queues/kafka"
	"goliath/queues/kafka/messages"
	"goliath/types/api"
)

type Log struct {}

func (_ Log) GetPath() string {
	return "/api/v1/log"
}

func (_ Log) GetMethod() string {
	return http.MethodPost
}

func (_ Log) DoHandle(c echo.Context) error {
	log := new(api.Log)
	if err := c.Bind(log); err != nil {
		return err
	}

	err := kafka.Send(messages.Log{Level: log.Level, Message: log.Message})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
