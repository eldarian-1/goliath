package messages

type Message interface {
	GetTopic() string
	ToBytes() ([]byte, error)
}
