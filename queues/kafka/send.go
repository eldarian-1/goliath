package kafka

import (
	"context"
	"errors"

	"github.com/segmentio/kafka-go"

	"goliath/queues/kafka/consumers"
	"goliath/queues/kafka/messages"
)

type Reader struct {
	consumer consumers.Consumer
	kafkaReader *kafka.Reader
}

const kafkaHost = "localhost:9092"

var kafkaWriterMap map[string]*kafka.Writer
var readerMap map[string]Reader

func init() {
	supportedMessages := []messages.Message{
		messages.Log{},
	}
	supportedConsumers := []consumers.Consumer{
		consumers.Log{},
	}
	kafkaWriterMap = make(map[string]*kafka.Writer)
	readerMap = make(map[string]Reader)

	for _, message := range supportedMessages {
		topic := message.GetTopic()
		writer := kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{kafkaHost},
			Topic:   topic,
		})
		defer writer.Close()
		kafkaWriterMap[topic] = writer
	}

	for _, consumer := range supportedConsumers {
		topic := consumer.GetTopic()
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{kafkaHost},
			Topic:   topic,
		})
		defer reader.Close()
		readerMap[topic] = Reader{
			consumer: consumer,
			kafkaReader: reader,
		}
	}

	startTopicsProcessing()
}

func Send(message messages.Message) error {
	writer, ok := kafkaWriterMap[message.GetTopic()]
	if !ok {
		return errors.New("Message topic is not supported")
	}

	return writer.WriteMessages(context.Background(), kafka.Message{
		Value: message.ToBytes(),
	})
}

func startTopicsProcessing() {
	for _, reader := range readerMap {
		go processTopic(reader)
	}
}

func processTopic(reader Reader) {
	for {
		msg, err := reader.kafkaReader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}

		err = reader.consumer.Process(msg.Value)
		if err != nil {
			panic(err)
		}

		err = reader.kafkaReader.CommitMessages(context.Background(), msg)
		if err != nil {
			panic(err)
		}
	}
}
