package messagequeue

import (
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"log"
	errorhandler "picpay-challenge-go/cmd/api/error"
)

type RabbitMQAdapter struct {
	Amqp      *amqp091.Channel
	QueueName string
}

func (adapter *RabbitMQAdapter) Publish(amount float64, payerId int, receiverId int) errorhandler.APIError {
	event := TransferMoneyEvent{amount, payerId, receiverId}

	jsonEvent, err := json.Marshal(&event)

	if err != nil {
		return nil
	}

	result := adapter.Amqp.Publish("", adapter.QueueName, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        jsonEvent,
	})

	if result != nil {
		log.Printf("Error publishing message to queue: %s", result.Error())
	}

	return nil
}

func (a *RabbitMQAdapter) Consume() TransferMoneyEvent {
	var event TransferMoneyEvent

	messages, err := a.Amqp.Consume(a.QueueName, "", true, false, false, false, nil)

	if err != nil {
		log.Printf("Error consuming message from queue: %s", err)
	}

	for message := range messages {
		err := json.Unmarshal(message.Body, &event)

		if err != nil {
			log.Printf("Error unmarshalling message from queue: %s", err)
		}

		return event
	}

	return event
}
