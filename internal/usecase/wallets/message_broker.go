package walletusecase

import (
	errorhandler "picpay-challenge-go/cmd/api/error"
	"picpay-challenge-go/pkg/messagequeue"
)

type MessageBroker interface {
	Publish(amount float64, payerId int, receiverId int) errorhandler.APIError
	Consume() messagequeue.TransferMoneyEvent
}
