package messagequeue

type TransferMoneyEvent struct {
	Value      float64 `json:"value" validate:"required"`
	PayerId    int     `json:"payer_id" validate:"required"`
	ReceiverId int     `json:"receiver_id" validate:"required"`
}
