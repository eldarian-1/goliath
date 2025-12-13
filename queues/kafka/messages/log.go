package messages

type Log struct {
	Level	string `json:"level"`
	Message string `json:"message"`
}

func (_ Log) GetTopic() string {
	return "log"
}

func (_ Log) ToBytes() []byte {
	return []byte("log")
}
