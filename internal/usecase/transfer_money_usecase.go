package usecase

import (
	"picpay-challenge-go/cmd/api/dtos"
	errorhandler "picpay-challenge-go/cmd/api/error"
	"picpay-challenge-go/internal/domain"
	"picpay-challenge-go/internal/usecase/errors"
)

type TransferMoneyUseCase struct {
	Repository WalletRepository
}

func (wallet *TransferMoneyUseCase) Execute(dto dtos.TransferMoneyDTO) (domain.Wallet, errorhandler.APIError) {
	payerWallet, err := wallet.Repository.FindByUserId(dto.Payer)

	if err != nil {
		return domain.Wallet{}, err
	}

	if !payerWallet.HasBalance(dto.Value) {
		return domain.Wallet{}, &errors.InsufficientBalanceError{UserID: dto.Payer, Balance: dto.Value}
	}

	return payerWallet, nil
}
