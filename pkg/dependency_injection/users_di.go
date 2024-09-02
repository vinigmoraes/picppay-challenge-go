package dependency_injection

import (
	"gorm.io/gorm"
	"picpay-challenge-go/cmd/api/handlers"
	"picpay-challenge-go/internal/usecase"
	dbusers "picpay-challenge-go/pkg/database/users"
	dbwallets "picpay-challenge-go/pkg/database/wallet"
)

func InjectCreateUserHandler(db *gorm.DB) handlers.CreateUserHandler {
	return handlers.CreateUserHandler{
		UseCase: usecase.CreateUserUseCase{
			UserRepository:   &dbusers.UserRepositoryAdapter{DB: db},
			WalletRepository: &dbwallets.WalletRepositoryAdapter{DB: db},
		}}
}

func InjectTransferMoneyHandler(db *gorm.DB) handlers.TransferMoneyHandler {
	return handlers.TransferMoneyHandler{
		UseCase: usecase.TransferMoneyUseCase{},
	}
}
