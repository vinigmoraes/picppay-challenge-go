package usecase

import (
	errorhandler "picpay-challenge-go/cmd/api/error"
	"picpay-challenge-go/internal/domain"
)

type UserRepository interface {
	Save(model *domain.User) errorhandler.APIError
}
