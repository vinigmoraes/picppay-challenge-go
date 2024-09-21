package authorizer

type TransactionAuthorizerResponse struct {
	Status string `json:"status"`
	Data   Data
}

type Data struct {
	Authorization bool `json:"authorization"`
}

func (t TransactionAuthorizerResponse) IsAuthorized() bool {
	return t.Status == "success"
}
