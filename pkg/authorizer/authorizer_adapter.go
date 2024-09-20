package authorizer

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"picpay-challenge-go/pkg/httpclient"
)

type TransferMoneyAuthorizerAdapter struct {
	AuthorizerURL string
}

func (t TransferMoneyAuthorizerAdapter) Authorize() (TransactionAuthorizerResponse, error) {
	authorizerResponse := TransactionAuthorizerResponse{}
	uri, _ := url.Parse(t.AuthorizerURL)

	body, err := httpclient.Request(http.MethodGet, uri.String())

	if err != nil {
		log.Fatal(err)
	}

	jsonErr := json.Unmarshal(body, &authorizerResponse)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return authorizerResponse, nil
}
