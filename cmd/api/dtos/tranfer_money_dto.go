package dtos

type TransferMoneyDTO struct {
	Value    float64 `json:"value" validate:"required"`
	Payer    int     `json:"payer" validate:"required"`
	Receiver int     `json:"receiver" validate:"required"`
}
