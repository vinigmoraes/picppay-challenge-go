package walletusecase

import (
	"log"
	"picpay-challenge-go/pkg/messagequeue"
)

type PutMoneyInWalletUseCase struct {
	WalletRepository WalletRepository
	Authorizer       Authorizer
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

	if authorize.IsAuthorized() {
		receiverWallet.AddBalance(event.Value)
		p.WalletRepository.Update(&receiverWallet)
	}
}
