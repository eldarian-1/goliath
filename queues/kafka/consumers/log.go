package consumers

type Log struct {}

func (_ Log) GetTopic() string {
	return "log"
}

func (_ Log) Process(message []byte) error {
	return nil
}
