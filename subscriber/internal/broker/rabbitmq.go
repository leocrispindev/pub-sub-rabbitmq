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

	connTimeout := 5 * time.Second

	dialer := &net.Dialer{
		Timeout: connTimeout,
	}

	amqpConfig := amqp.Config{
		Dial: func(network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
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

}

func GetChannel() *amqp.Channel {
	return channel
}
