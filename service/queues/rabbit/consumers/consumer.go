package consumers

type Consumer interface {
	GetQueue() string
	Process(message []byte) error
}
