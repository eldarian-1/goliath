package main

import (
	"context"
	"fmt"

	"goliath/handlers"
	"goliath/migrations"
	"goliath/queues/kafka"
	"goliath/queues/rabbit"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := migrations.Migrate(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rabbit.StartRabbitConsumers(ctx)
	kafka.StartKafkaConsumers(ctx)
	handlers.Define()
}
