package usersusecase

import (
	"picpay-challenge-go/cmd/api/dtos"
	errorhandler "picpay-challenge-go/cmd/api/error"
	"picpay-challenge-go/internal/domain"
	usecase "picpay-challenge-go/internal/usecase/wallets"
)

const walletInitialBalance = 100.0

type CreateUserUseCase struct {
	UserRepository   UserRepository
	WalletRepository usecase.WalletRepository
}

func (useCase *CreateUserUseCase) Execute(userDTO dtos.UserDTO) (domain.User, errorhandler.APIError) {
	user := domain.CreateUser(userDTO)

	err := useCase.UserRepository.Save(&user)

	wallet := domain.Wallet{UserID: user.ID, Balance: walletInitialBalance}

	useCase.WalletRepository.Save(&wallet)

	return user, err
}
