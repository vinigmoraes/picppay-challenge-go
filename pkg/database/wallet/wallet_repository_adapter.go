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

func (adapter *WalletRepositoryAdapter) FindByUserId(userId int) (wallet domain.Wallet, apiError errorhandler.APIError) {
	wallets := Wallets{}

	result := adapter.DB.Find(wallets, userId)

	if result.Error != nil {
		return domain.Wallet{}, handleDatabaseError(wallets, result.Error)
	}

	result.Model(&wallets)

	return domain.Wallet{ID: wallets.ID, UserID: wallets.UserID, Balance: wallets.Balance}, nil
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

func (adapter *WalletRepositoryAdapter) Update(wallet *domain.Wallet) (apiError errorhandler.APIError) {
	var dbError errorhandler.APIError

	return dbError
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
