package main

import (
	"publisher-subscriber-rabbitmq/publisher/internal/broker"
	dispatcher "publisher-subscriber-rabbitmq/publisher/internal/services/dispatcher"
	"time"
)

func main() {
	timeToConnect := 5 * time.Second
	time.Sleep(timeToConnect)

	broker.Init()
	dispatcher.Init()
}
