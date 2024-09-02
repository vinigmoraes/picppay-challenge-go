package dberrors

import (
	"fmt"
	errorhandler "picpay-challenge-go/cmd/api/error"
)

type WalletAlreadyExistsError struct {
	WalletID int
}

func (e *WalletAlreadyExistsError) GetError() string {
	return fmt.Sprintf("Wallet already exist for user: %s", e.WalletID)
}

func (e *WalletAlreadyExistsError) GetErrorType() errorhandler.APIError {
	return e
}
