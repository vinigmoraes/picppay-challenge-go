package dbusers

import (
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	errorhandler "picpay-challenge-go/cmd/api/error"
	"picpay-challenge-go/internal/domain"
	dberrors "picpay-challenge-go/pkg/database/errors"
	"time"
)

type UserRepositoryAdapter struct {
	DB *gorm.DB
}

func (adapter *UserRepositoryAdapter) Save(user *domain.User) errorhandler.APIError {
	model := Users{
		Name:                  user.Name,
		Email:                 user.Email,
		Password:              user.Password,
		CPF:                   user.CPF,
		IsAbleToTransferMoney: user.IsAbleToTransferMoney(),
		Status:                "active",
		CreatedAt:             time.Now(),
		Type:                  user.Type.String(),
	}

	result := adapter.DB.Create(&model)

	user.SetId(model.ID)

	return handleDatabaseError(model, result.Error)
}

func handleDatabaseError(user Users, err error) errorhandler.APIError {
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
