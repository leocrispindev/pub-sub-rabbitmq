package main

import (
	"publisher-subscriber-rabbitmq/publisher/internal/broker"
	dispatcher "publisher-subscriber-rabbitmq/publisher/internal/services/dispatcher"
)

func main() {
	broker.Init()
	dispatcher.Init()
}
