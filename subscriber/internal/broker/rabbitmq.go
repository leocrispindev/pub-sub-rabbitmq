package broker

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var channel *amqp.Channel

func Init() {
	initConnection()
}

func initConnection() {
	conn, err := amqp.Dial("amqp://userguest:user123@localhost:5672/")

	if err != nil {
		log.Panic("Error on open connection to broker", err)
	}

	log.Println("Connection OK")
	openChannel(conn)

}

func openChannel(conn *amqp.Connection) {
	ch, err := conn.Channel()

	if err != nil {
		log.Panic("Error to open channel with broker")

	}

	log.Println("Channel OK")
	channel = ch

}

func GetChannel() *amqp.Channel {
	return channel
}
