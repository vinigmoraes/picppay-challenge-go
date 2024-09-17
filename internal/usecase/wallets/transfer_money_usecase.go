package walletusecase

import (
	"picpay-challenge-go/cmd/api/dtos"
	errorhandler "picpay-challenge-go/cmd/api/error"
	"picpay-challenge-go/internal/domain"
	"picpay-challenge-go/internal/usecase/errors"
)

type TransferMoneyUseCase struct {
	WalletRepository WalletRepository
	Publisher        MessageBroker
}

func (w TransferMoneyUseCase) Execute(dto dtos.TransferMoneyDTO) (domain.Wallet, errorhandler.APIError) {
	payerWallet, err := w.WalletRepository.FindByUserId(dto.Payer)

	if err != nil {
		return domain.Wallet{}, err
	}

	if !payerWallet.HasBalance(dto.Value) {
		return domain.Wallet{}, &errors.InsufficientBalanceError{UserID: dto.Payer, Balance: dto.Value}
	}

	payerWallet.RemoveBalance(dto.Value)

	w.WalletRepository.Update(&payerWallet)

	w.Publisher.Publish(dto.Value, dto.Payer, dto.Receiver)

	return payerWallet, nil
}
