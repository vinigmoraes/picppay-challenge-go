package consumers

import (
	"encoding/json"
	"log"
	walletusecase "picpay-challenge-go/internal/usecase/wallets"
	"picpay-challenge-go/pkg/messagequeue"
)

type GiveMoneyBackConsumer struct {
	Broker  walletusecase.MessageBroker
	UseCase walletusecase.GiveMoneyBackWalletUseCase
}

func (g GiveMoneyBackConsumer) Consume() {
	go func() {
		forever := make(chan bool)
		event := messagequeue.TransferMoneyEvent{}

		messages := g.Broker.Consume()

		for message := range messages {
			err := json.Unmarshal(message.Body, &event)

			if err != nil {
				log.Printf("Error unmarshalling message from queue: %s", err)
			}

			g.UseCase.Execute(event)
		}
		<-forever
	}()
}
