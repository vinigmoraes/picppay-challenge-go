package walletusecase

import (
	"log"
	"log/slog"
	"picpay-challenge-go/pkg/messagequeue"
)

type PutMoneyInWalletUseCase struct {
	WalletRepository WalletRepository
	Authorizer       Authorizer
	MessageBroker    MessageBroker
}

func (p PutMoneyInWalletUseCase) Execute(event messagequeue.TransferMoneyEvent) {
	receiverWallet, dbErr := p.WalletRepository.FindByUserId(event.ReceiverId)

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	authorize, authorizeErr := p.Authorizer.Authorize()

	if authorizeErr != nil {
		log.Fatal(authorizeErr)
	}

	if !authorize.IsAuthorized() {
		slog.Warn("Transaction not authorized for payer", slog.Int("payer_id", event.PayerId))
		p.MessageBroker.Publish(event.Value, event.PayerId, event.ReceiverId)
		return
	}

	receiverWallet.AddBalance(event.Value)
	p.WalletRepository.Update(&receiverWallet)
}
