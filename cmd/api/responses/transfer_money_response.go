package responses

type TransferMoneyResponse struct {
	UserID  int     `json:"user_id"`
	Balance float64 `json:"balance"`
}
