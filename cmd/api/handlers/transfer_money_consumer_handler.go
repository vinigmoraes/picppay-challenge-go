package handlers

import walletusecase "picpay-challenge-go/internal/usecase/wallets"

type TransferMoneyConsumerHandler struct {
	Broker  walletusecase.MessageBroker
	UseCase walletusecase.PutMoneyInWalletUseCase
}

func (t *TransferMoneyConsumerHandler) Consume() {
	go func() {
		forever := make(chan bool)

		event := t.Broker.Consume()
		t.UseCase.Execute(event)

		<-forever
	}()
}
