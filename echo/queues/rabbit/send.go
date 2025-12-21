package rabbit

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"goliath/queues/rabbit/messages"
)

func Send(message messages.Message) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		message.GetQueue(), // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := message.ToBytes()
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(
		ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: message.GetContentType(),
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	log.Printf(" [x] Sent %s\n", body)

	return nil
}
