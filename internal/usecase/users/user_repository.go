package usecase

import dbusers "picpay-challenge-go/pkg/database/users"

type UserRepository interface {
	Save(model dbusers.UserModel) error
}
