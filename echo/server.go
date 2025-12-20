package main

import (
	"goliath/handlers"
	"goliath/migrations"
	"goliath/queues/kafka"
)

func main() {
	defer kafka.CloseAllConnections()

	migrations.Migrate()
	handlers.Define()
}
