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

	defer conn.Close()

}

func openChannel(conn *amqp.Connection) {
	ch, err := conn.Channel()

	if err != nil {
		log.Panic("Error to open channel with broker")

	}

	log.Println("Channel OK")

	declareExchnage(ch)
}

// Criando e definindo os topicos de direcionamento na exchange
func declareExchnage(channel *amqp.Channel) {
	topics := []string{"logs_topic", "messages_topic"}

	//Interando sobre a lista de topicos que a exchnage mapear
	for _, topic := range topics {
		err := channel.ExchangeDeclare(
			topic, "topic", true, false, false, false, nil, // arguments
		)

		if err != nil {
			log.Panic("Error to declare excnage topic: " + topic)

		}

		log.Println("Exhnage declared topic: " + topic)

	}

}

func GetChannel() *amqp.Channel {
	return channel
}
