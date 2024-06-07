package publish

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"publisher-subscriber-rabbitmq/publisher/internal/broker"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	broker *amqp.Channel
}

func (p *Publisher) Init() {
	p.broker = broker.GetChannel()
}

func (p *Publisher) Send(ctx context.Context, topic string, body []byte) error {
	err := p.broker.PublishWithContext(ctx,
		"example-rabbitmq_exchange", // topico mapeado na exhcnage
		topic,                       // a chave da mensagem é gerada a partir do body enviado - objetivo é evitar chave duplicada
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	return err
}

func generateKey(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}
