package walletusecase

import (
	"github.com/rabbitmq/amqp091-go"
	errorhandler "picpay-challenge-go/cmd/api/error"
)

type MessageBroker interface {
	Publish(amount float64, payerId int, receiverId int) errorhandler.APIError
	Consume() <-chan amqp091.Delivery
}
