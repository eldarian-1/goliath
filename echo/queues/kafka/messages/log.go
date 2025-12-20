package messages

import (
	"encoding/json"
)

type Log struct {
	Level	string `json:"level"`
	Message string `json:"message"`
}

func (_ Log) GetTopic() string {
	return "log"
}

func (l Log) ToBytes() ([]byte, error) {
	return json.Marshal(l)
}
