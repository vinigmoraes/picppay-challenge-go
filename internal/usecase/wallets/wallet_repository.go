package walletusecase

import (
	errorhandler "picpay-challenge-go/cmd/api/error"
	"picpay-challenge-go/internal/domain"
)

type WalletRepository interface {
	FindByUserId(userId int) (wallet domain.Wallet, apiError errorhandler.APIError)
	Save(wallet *domain.Wallet) (apiError errorhandler.APIError)
	Update(wallet *domain.Wallet) (apiError errorhandler.APIError)
}
