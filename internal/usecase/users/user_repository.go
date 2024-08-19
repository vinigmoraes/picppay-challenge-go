package usecase

import (
	errorhandler "picpay-challenge-go/cmd/api/error"
	dbusers "picpay-challenge-go/pkg/database/users"
)

type UserRepository interface {
	Save(model *dbusers.Users) errorhandler.APIError
}
