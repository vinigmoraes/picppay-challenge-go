package usecase

import (
	"picpay-challenge-go/cmd/api/dtos"
	database "picpay-challenge-go/pkg/database/users"
	"time"
)

type CreateUserUseCase struct {
	Repository UserRepository
}

func (useCase *CreateUserUseCase) Execute(dto dtos.UserDTO) {
	model := database.Users{
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  dto.Password,
		CPF:       dto.CPF,
		Status:    "active",
		CreatedAt: time.Now(),
	}

	err := useCase.Repository.Save(&model)

	if err != nil {
		return
	}
}
