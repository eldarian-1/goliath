package messages

type Message interface {
	GetQueue() string
	GetContentType() string
	ToBytes() ([]byte, error)
}
