package walletusecase

import "picpay-challenge-go/pkg/authorizer"

type Authorizer interface {
	Authorize() (authorizer.TransactionAuthorizerResponse, error)
}
