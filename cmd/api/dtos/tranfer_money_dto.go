package dtos

type TransferMoneyDTO struct {
	Value    float64 `json:"value" validate:"required"`
	Payer    int     `json:"payer,string" validate:"required"`
	Receiver int     `json:"receiver,string" validate:"required"`
}
