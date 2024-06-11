package dispatch

import (
	"context"
	"fmt"
	pubService "publisher-subscriber-rabbitmq/publisher/internal/services/publisher"
	"time"

	"github.com/google/uuid"
)

var publisher *pubService.Publisher

func Init() {
	publisher = new(pubService.Publisher)
	publisher.Init()

	startRandomMessages()
}

func startRandomMessages() {
	ctx := context.Background()

	log(ctx, "INFO", "Iniciando dispatch")

	for {
		proccessId := uuid.New().String()

		log(ctx, "INFO", "Start dispatch batch messages [id]= "+proccessId)

		for i := 0; i < 10; i++ {
			msg, id := generateMessage()

			err := publisher.Send(ctx, "messages_topic.key", []byte(msg))

			if err != nil {
				println(err.Error())
				log(ctx, "ERROR", fmt.Sprintf("Publish error message [id]=%s", id))
				continue
			}

			log(ctx, "INFO", fmt.Sprintf("Publish success message [id]=%s", id))

		}

		log(ctx, "INFO", "Finish dispatch batch messages [id]= "+proccessId)
		time.Sleep(2 * time.Second)

	}

}

func generateMessage() (string, string) {
	messageId := uuid.New().String()

	message := fmt.Sprintf("Mensagem de teste de publisher/subscriber com RabbitMq [id]=%s", messageId)

	return message, messageId
}

func log(ctx context.Context, level string, message string) {
	body := fmt.Sprintf("[LEVEL]=%s [BODY]=%s", level, message)

	publisher.Send(ctx, "logs_topic.key", []byte(body))
}
