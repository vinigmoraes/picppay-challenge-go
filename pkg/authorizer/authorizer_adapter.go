package authorizer

import (
	"encoding/json"
	"log"
	"net/http"
	"picpay-challenge-go/pkg/httpclient"
)

type TransferMoneyAuthorizerAdapter struct {
	AuthorizerURL string
}

func (t TransferMoneyAuthorizerAdapter) Authorize() (TransactionAuthorizerResponse, error) {
	authorizerResponse := TransactionAuthorizerResponse{}

	body, err := httpclient.Request(http.MethodGet, t.AuthorizerURL)

	if err != nil {
		log.Fatal(err)
	}

	jsonErr := json.Unmarshal(body, &authorizerResponse)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return authorizerResponse, nil
}
