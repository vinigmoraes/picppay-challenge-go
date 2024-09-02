package dbwallets

import (
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	errorhandler "picpay-challenge-go/cmd/api/error"
	"picpay-challenge-go/internal/domain"
	dberrors "picpay-challenge-go/pkg/database/errors"
)

type WalletRepositoryAdapter struct {
	DB *gorm.DB
}

func (adapter *WalletRepositoryAdapter) FindByUserId(userId string) (wallet domain.Wallet, apiError errorhandler.APIError) {
	return domain.Wallet{}, nil
}

func (adapter *WalletRepositoryAdapter) Save(wallet *domain.Wallet) errorhandler.APIError {
	wallets := Wallets{
		UserID:  wallet.UserID,
		Balance: wallet.Balance,
	}

	result := adapter.DB.Create(&wallets)

	if result.Error != nil {
		return handleDatabaseError(wallets, result.Error)
	}

	wallet.SetId(wallets.ID)

	return nil
}

func handleDatabaseError(wallets Wallets, err error) errorhandler.APIError {
	var dbError errorhandler.APIError

	if err != nil {
		pgerror := err.(*pgconn.PgError)

		switch pgerror.Code {
		case "23505":
			dbError = &dberrors.WalletAlreadyExistsError{WalletID: wallets.UserID}
		}
	}
	return dbError
}
