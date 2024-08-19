package dbusers

import (
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	errorhandler "picpay-challenge-go/cmd/api/error"
	dberrors "picpay-challenge-go/pkg/database/errors"
)

type UserRepositoryAdapter struct {
	DB *gorm.DB
}

func (adapter *UserRepositoryAdapter) Save(model *Users) errorhandler.APIError {
	result := adapter.DB.Create(&model)

	return handleDatabaseError(model, result.Error)
}

func handleDatabaseError(user *Users, err error) errorhandler.APIError {
	var dbError errorhandler.APIError

	if err != nil {
		pgerror := err.(*pgconn.PgError)

		switch pgerror.Code {
		case "23505":
			dbError = &dberrors.UserAlreadyExistsError{CPF: user.CPF}
		}
	}
	return dbError
}
