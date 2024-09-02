package usecase

import (
	errorhandler "picpay-challenge-go/cmd/api/error"
	"picpay-challenge-go/internal/domain"
)

type WalletRepository interface {
	FindByUserId(userId string) (wallet domain.Wallet, apiError errorhandler.APIError)
	Save(wallet *domain.Wallet) (apiError errorhandler.APIError)
}
