package dberrors

import (
	"fmt"
	errorhandler "picpay-challenge-go/cmd/api/error"
)

type UserAlreadyExistsError struct {
	CPF string
}

func (e *UserAlreadyExistsError) GetError() string {
	return fmt.Sprintf("%s user already exists", e.CPF)
}

func (e *UserAlreadyExistsError) GetErrorType() errorhandler.APIError {
	return e
}
