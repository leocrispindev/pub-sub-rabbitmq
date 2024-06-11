package broker

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var channel *amqp.Channel

func Init() {
	initConnection()
}

func initConnection() {

	host := os.Getenv("BROKER_HOST")

	if host == "" {
		host = "localhost"
	}

	connectionHost := fmt.Sprintf("amqp://userguest:user123@%s:5672/", host)

	connTimeout := 10 * time.Second

	amqpConfig := amqp.Config{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, connTimeout)
		},
	}

	conn, err := amqp.DialConfig(connectionHost, amqpConfig)

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

	declareExchange(channel)
	bindQueues(channel)

}

func declareExchange(channel *amqp.Channel) {

	err := channel.ExchangeDeclare(
		"example-rabbitmq_exchange", "topic", true, false, false, false, nil, // arguments
	)

	if err != nil {
		log.Panic("Error on declare exchange")
	}

	log.Println("Success exchange declare")

}

// Esse fluxo só é necessário para que a aplicação crie as filas de mensagens na hora que inicializar
func bindQueues(ch *amqp.Channel) {
	// Declarar uma fila
	qMessages, err := ch.QueueDeclare(
		"queue_messages", // nome da fila
		true,             // durable - para manter durante o restart do rabbitmq
		false, false, false, nil,
	)

	if err != nil {
		log.Panic("Error on create queue", err)
	}

	// Cria a 'conexão' entre a exchange e a fila
	err = ch.QueueBind(
		qMessages.Name,              // nome da fila
		"messages_topic.key",        // chave de roteamento
		"example-rabbitmq_exchange", // nome da exchange
		false,
		nil,
	)

	if err != nil {
		log.Panic("Error on bind queue", err)
	}

	log.Printf("Created and bind [queue]=%s", qMessages.Name)

	qLogs, err := ch.QueueDeclare(
		"queue_logs", // nome da fila
		true,         // durable - para manter durante o restart do rabbitmq
		false, false, false, nil,
	)

	err = ch.QueueBind(
		qLogs.Name,
		"logs_topic.key",
		"example-rabbitmq_exchange",
		false,
		nil,
	)

	log.Printf("Created and bind [queue]=%s", qMessages.Name)

}

func GetChannel() *amqp.Channel {
	return channel
}
