package messagequeue

import (
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"log/slog"
	errorhandler "picpay-challenge-go/cmd/api/error"
)

type RabbitMQAdapter struct {
	Amqp      *amqp091.Channel
	QueueName string
}

func (a RabbitMQAdapter) Publish(amount float64, payerId int, receiverId int) errorhandler.APIError {
	event := TransferMoneyEvent{amount, payerId, receiverId}

	jsonEvent, err := json.Marshal(&event)

	if err != nil {
		return nil
	}

	result := a.Amqp.Publish("", a.QueueName, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        jsonEvent,
	})

	if result != nil {
		log.Printf("Error publishing message to queue: %s", result.Error())
	}

	slog.Info("Message published successfully", slog.String("queue_name", a.QueueName))

	return nil
}

func (a RabbitMQAdapter) Consume() <-chan amqp091.Delivery {
	messages, err := a.Amqp.Consume(a.QueueName, "", false, false, false, false, nil)

	if err != nil {
		log.Printf("Error consuming message from queue: %s", err)
	}

	return messages
}
