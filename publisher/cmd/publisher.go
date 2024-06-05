package main

import (
	"publisher-subscriber-rabbitmq/publisher/internal/broker"
	"publisher-subscriber-rabbitmq/publisher/internal/services/dispatch"
)

func main() {
	broker.Init()
	dispatch.Init()
}
