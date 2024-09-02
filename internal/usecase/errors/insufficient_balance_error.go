package errors

import (
	"fmt"
	errorhandler "picpay-challenge-go/cmd/api/error"
)

type InsufficientBalanceError struct {
	UserID  string
	Balance float64
}

func (e *InsufficientBalanceError) GetError() string {
	return fmt.Sprintf("User: %s has insufficient balance, currently balance is: %s", e.UserID, e.Balance)
}

func (e *InsufficientBalanceError) GetErrorType() errorhandler.APIError {
	return e
}
