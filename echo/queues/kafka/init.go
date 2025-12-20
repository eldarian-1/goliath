package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"

	"goliath/queues/kafka/consumers"
	"goliath/queues/kafka/messages"
)

type reader struct {
	consumer consumers.Consumer
	kafkaReader *kafka.Reader
}

const kafkaHost = "localhost:9092"

var kafkaWriterMap map[string]*kafka.Writer
var readerMap map[string]reader

func init() {
	initWriters()
	initReaders()
	startTopicsProcessing()
}

func initWriters() {
	supportedMessages := []messages.Message{
		messages.Log{},
	}
	kafkaWriterMap = make(map[string]*kafka.Writer)

	for _, message := range supportedMessages {
		topic := message.GetTopic()
		writer := kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{kafkaHost},
			Topic:   topic,
		})
		kafkaWriterMap[topic] = writer
	}
}

func initReaders() {
	supportedConsumers := []consumers.Consumer{
		consumers.Log{},
	}
	readerMap = make(map[string]reader)

	for _, consumer := range supportedConsumers {
		topic := consumer.GetTopic()
		kafkaReader := kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{kafkaHost},
			Topic:   topic,
		})
		readerMap[topic] = reader{
			consumer: consumer,
			kafkaReader: kafkaReader,
		}
	}
}

func startTopicsProcessing() {
	for _, reader := range readerMap {
		go processTopic(reader)
	}
}

func processTopic(reader reader) {
	for {
		msg, err := reader.kafkaReader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}

		err = reader.consumer.Process(msg.Value)
		if err != nil {
			panic(err)
		}
	}
}

func CloseAllConnections() {
	fmt.Println("Close all Kafka coonections")
	if kafkaWriterMap != nil {
		for _, writer := range kafkaWriterMap {
			writer.Close()
		}
	}
	if readerMap != nil {
		for _, reader := range readerMap {
			reader.kafkaReader.Close()
		}
	}
}
