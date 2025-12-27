package rabbit

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"

	"goliath/queues/rabbit/consumers"
	"goliath/utils"
)

const (
	consumerName = "echo_server_rabbit_consumer"
)

var rabbitConnectionString string

func init() {
	rabbitConnectionString = fmt.Sprintf(
		"amqp://%s:%s@%s/",
		utils.GetEnv("RABBITMQ_USER", "guest"),
		utils.GetEnv("RABBITMQ_PASSWORD", "guest"),
		utils.GetEnv("RABBITMQ_HOST", "localhost:5672"),
	)
}

func StartRabbitConsumers(ctx context.Context) {
	consumers := []consumers.Consumer{
		consumers.Log{},
	}
	for _, consumer := range consumers {
		go startConsumer(ctx, consumer)
	}
}

func startConsumer(ctx context.Context, consumer consumers.Consumer) {
	conn, err := amqp.Dial(rabbitConnectionString)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		consumer.GetQueue(), // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,       // queue
		consumerName, // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")

	for {
		select {
		case d := <-msgs:
			err := consumer.Process(d.Body)
			if err != nil {
				log.Printf("Failed with error: %s\n", err.Error())
			}
		case <-ctx.Done():
			return
		}
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
