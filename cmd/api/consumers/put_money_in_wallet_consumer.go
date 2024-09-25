package consumers

import (
	"encoding/json"
	"log"
	walletusecase "picpay-challenge-go/internal/usecase/wallets"
	"picpay-challenge-go/pkg/messagequeue"
)

type TransferMoneyConsumerHandler struct {
	Broker  walletusecase.MessageBroker
	UseCase walletusecase.PutMoneyInWalletUseCase
}

func (t TransferMoneyConsumerHandler) Consume() {
	go func() {
		event := messagequeue.TransferMoneyEvent{}

		forever := make(chan bool)

		messages := t.Broker.Consume()

		for message := range messages {
			err := json.Unmarshal(message.Body, &event)

			if err != nil {
				log.Printf("Error unmarshalling message from queue: %s", err)
			}

			t.UseCase.Execute(event)
		}
		<-forever
	}()
}
