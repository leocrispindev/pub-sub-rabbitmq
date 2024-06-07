package consumer

import (
	"publisher-subscriber-rabbitmq/subscriber/internal/broker"

	amqp "github.com/rabbitmq/amqp091-go"
)

// o Consumidor vai receber uma funçao que recebe cada uma das mensagem e lida com ela do jeito que definir
type Consumer struct {
	broker      *amqp.Channel
	handlerFunc func(msg amqp.Delivery)
	queueName   string
}

func (c *Consumer) Constructor(function func(msg amqp.Delivery), queueName string) {
	c.broker = broker.GetChannel()
	c.handlerFunc = function
	c.queueName = queueName
}

func (c *Consumer) Start() error {
	// ao instancias um consumidor, é devolvido um channel que receberá as mensagens da fila
	chDeliveryMessages, err := c.broker.Consume(c.queueName, "consumer-id", true, false, false, false, nil)

	if err != nil {
		return err
	}

	go func() {
		// interando sobre as mensagens consumidar da fila
		// cada consumidor pode ter um funçao diferente
		for message := range chDeliveryMessages {
			c.handlerFunc(message)
		}
	}()

	return nil
}
