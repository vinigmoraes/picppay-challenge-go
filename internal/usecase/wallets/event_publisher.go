package walletusecase

import (
	errorhandler "picpay-challenge-go/cmd/api/error"
	"picpay-challenge-go/internal/domain"
)

type EventPublisher interface {
	Publish(wallet *domain.Wallet) errorhandler.APIError
}
