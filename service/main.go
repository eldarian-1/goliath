package main

import (
	"context"
	"fmt"

	"goliath/migrations"
	"goliath/queues/kafka"
	"goliath/server"
	// "goliath/queues/rabbit"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
			cancel()
		}
	}()

	err := migrations.Migrate(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// rabbit.StartRabbitConsumers(ctx)
	go kafka.StartKafkaConsumers(ctx)
	server.Define()
}
