package consumers

import (
	"encoding/json"
	"errors"
	"fmt"

	"goliath/queues/kafka/messages"
)

type Log struct {}

func (_ Log) GetTopic() string {
	return "log"
}

func (_ Log) Process(message []byte) error {
	var log messages.Log
	err := json.Unmarshal(message, &log)
	if err != nil {
		return errors.New("Deserializing of log was failed")
	}

	process(log)

	return nil
}

func process(log messages.Log) {
	fmt.Printf("%s: %s\n", log.Level, log.Message)
}
