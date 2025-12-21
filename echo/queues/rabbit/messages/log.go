package messages

import (
	"encoding/json"
)

type Log struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func (_ Log) GetQueue() string {
	return "log"
}

func (_ Log) GetContentType() string {
	return "application/json"
}

func (l Log) ToBytes() ([]byte, error) {
	return json.Marshal(l)
}
