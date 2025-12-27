package consumers

type Consumer interface {
	GetTopic() string
	Process(message []byte) error
}
