package dtos

type TransferMoneyDTO struct {
	Value    float64 `json:"value" validate:"required"`
	Payer    string  `json:"payer" validate:"required"`
	Receiver string  `json:"receiver" validate:"required"`
}
