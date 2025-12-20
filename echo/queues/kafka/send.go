package kafka

import (
	"context"
	"errors"

	"github.com/segmentio/kafka-go"

	"goliath/queues/kafka/messages"
)

func Send(message messages.Message) error {
	writer, ok := kafkaWriterMap[message.GetTopic()]
	if !ok {
		return errors.New("Message topic is not supported")
	}

	bytes, err := message.ToBytes()
	if err != nil {
		return errors.New("Message convertings to bytes has error")
	}

	return writer.WriteMessages(context.Background(), kafka.Message{
		Value: bytes,
	})
}
