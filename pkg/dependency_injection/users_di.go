package dependency_injection

import (
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
	"picpay-challenge-go/cmd/api/handlers"
	usersusecase "picpay-challenge-go/internal/usecase/users"
	walletusecase "picpay-challenge-go/internal/usecase/wallets"
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
			Publisher:        &messagequeue.EventPublisherAdapter{Amqp: amqp},
		},
	}
}
