package services

import (
	"github.com/segmentio/kafka-go"
)

type IDependencies interface {
	GetKafkaWriter() *kafka.Writer
}

var Dependencies IDependencies

type realDependencies struct {
	kafkaWriter *kafka.Writer
}

func (d realDependencies) GetKafkaWriter() *kafka.Writer {
	return d.kafkaWriter
}

func init() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "my-topic",
	})
	defer writer.Close()

	Dependencies = realDependencies{
		kafkaWriter: writer,
	}
}
