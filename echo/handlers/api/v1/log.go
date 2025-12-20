package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"goliath/queues/kafka"
	kafka_messages "goliath/queues/kafka/messages"
	rabbit_messages "goliath/queues/rabbit/messages"
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

	err := sendMessage(log)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}

func sendMessage(log *api.Log) error {
	if log.Broker != nil && *log.Broker == "rabbit" {
		return rabbit.Send(rabbit_messages.Log{Level: log.Level, Message: log.Message})
	}
	return kafka.Send(kafka_messages.Log{Level: log.Level, Message: log.Message})
}
