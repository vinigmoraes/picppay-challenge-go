package messagequeue

import (
	"github.com/rabbitmq/amqp091-go"
	errorhandler "picpay-challenge-go/cmd/api/error"
	"picpay-challenge-go/internal/domain"
)

type EventPublisherAdapter struct {
	Amqp      *amqp091.Channel
	QueueName string
}

func (adapter *EventPublisherAdapter) Publish(wallet *domain.Wallet) errorhandler.APIError {
	adapter.Amqp.Publish("", adapter.QueueName, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        []byte("{}"),
	})

	return nil
}
