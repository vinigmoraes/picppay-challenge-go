package walletusecase

import (
	"log/slog"
	"picpay-challenge-go/pkg/messagequeue"
)

type GiveMoneyBackWalletUseCase struct {
	WalletRepository WalletRepository
}

func (g GiveMoneyBackWalletUseCase) Execute(event messagequeue.TransferMoneyEvent) {
	payerWallet, dbError := g.WalletRepository.FindByUserId(event.PayerId)

	if dbError != nil {
		slog.Warn("Wallet not found for user", slog.Int("payer_id", event.PayerId))
		return
	}

	payerWallet.AddBalance(event.Value)
	g.WalletRepository.Update(&payerWallet)

	slog.Warn("Gave money back to payer wallet successfully", slog.Int("payer_id", event.PayerId))
}
