package main

import (
	"log"
	"os"
	"publisher-subscriber-rabbitmq/subscriber/internal/broker"
	csm "publisher-subscriber-rabbitmq/subscriber/internal/services/consumer"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	if len(os.Args) < 1 {
		panic("Invalid parameter: queue is missing")
	}

	broker.Init()

	consumer := new(csm.Consumer)

	consumer.Constructor(handleMessage, os.Args[1])

	err := consumer.Start()

	if err != nil {
		panic("Error on starting consumer " + err.Error())
	}

	var keepAlive chan string

	log.Printf(" Keeping consumer alive. To exit press CTRL+C")

	<-keepAlive
}

func handleMessage(message amqp091.Delivery) {
	messageId := message.MessageId
	textMessage := string(message.Body) // Deserialize o body da mensagem

	log.Printf("[id]=%s [body]=%s", messageId, textMessage)

}
