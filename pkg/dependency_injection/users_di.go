package dependency_injection

import (
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
	"picpay-challenge-go/cmd/api/handlers"
	usersusecase "picpay-challenge-go/internal/usecase/users"
	walletusecase "picpay-challenge-go/internal/usecase/wallets"
	"picpay-challenge-go/pkg/authorizer"
	dbusers "picpay-challenge-go/pkg/database/users"
	dbwallets "picpay-challenge-go/pkg/database/wallet"
	"picpay-challenge-go/pkg/messagequeue"
)

func InjectCreateUserHandler(db *gorm.DB) handlers.CreateUserHandler {
	return handlers.CreateUserHandler{
		UseCase: usersusecase.CreateUserUseCase{
			UserRepository:   &dbusers.UserRepositoryAdapter{DB: db},
			WalletRepository: &dbwallets.WalletRepositoryAdapter{DB: db},
		}}
}

func InjectTransferMoneyHandler(db *gorm.DB, amqp *amqp091.Channel) handlers.TransferMoneyHandler {
	return handlers.TransferMoneyHandler{
		UseCase: walletusecase.TransferMoneyUseCase{
			WalletRepository: &dbwallets.WalletRepositoryAdapter{DB: db},
			Publisher:        &messagequeue.RabbitMQAdapter{Amqp: amqp, QueueName: "transfer-money"},
		},
	}
}

func InjectTransferMoneyConsumer(db *gorm.DB, amqp *amqp091.Channel) handlers.TransferMoneyConsumerHandler {
	return handlers.TransferMoneyConsumerHandler{
		Broker: &messagequeue.RabbitMQAdapter{Amqp: amqp, QueueName: "transfer-money"},
		UseCase: walletusecase.PutMoneyInWalletUseCase{
			WalletRepository: &dbwallets.WalletRepositoryAdapter{DB: db},
			Authorizer:       authorizer.TransferMoneyAuthorizerAdapter{AuthorizerURL: "https://util.devi.tools/api/v2/authorize"},
			MessageBroker:    &messagequeue.RabbitMQAdapter{Amqp: amqp, QueueName: "transfer-money-dlq"},
		},
	}
}
