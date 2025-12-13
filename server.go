package main

import (
	"goliath/handlers"
	"goliath/migrations"
)

func main() {
	migrations.Migrate()
	handlers.Define()
}
