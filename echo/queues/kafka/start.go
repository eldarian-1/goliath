package kafka

import (
	"context"
	"sync"

	"github.com/segmentio/kafka-go"

	"goliath/queues/kafka/consumers"
	"goliath/queues/kafka/messages"
)

type reader struct {
	consumer    consumers.Consumer
	kafkaReader *kafka.Reader
}

const kafkaHost = "localhost:9092"

var kafkaWriterMap map[string]*kafka.Writer
var readerMap map[string]reader

func init() {
	initWriters()
	initReaders()
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
			consumer:    consumer,
			kafkaReader: kafkaReader,
		}
	}
}

func StartKafkaConsumers(ctx context.Context) {
	var wg sync.WaitGroup

	for _, reader := range readerMap {
		wg.Add(1)
		go func() {
			defer wg.Done()
			go processTopic(ctx, reader)
		}()
	}

	wg.Wait()
}

func processTopic(ctx context.Context, reader reader) {
	for {
		msg, err := reader.kafkaReader.ReadMessage(ctx)
		if err != nil {
			return
		}

		err = reader.consumer.Process(msg.Value)
		if err != nil {
			return
		}
	}
}
