package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picpay-challenge-go/cmd/api/dtos"
	"picpay-challenge-go/internal/usecase/users"
)

type CreateUserHandler struct {
	UseCase usecase.CreateUserUseCase
}

func (handler *CreateUserHandler) CreateUser(ctx *gin.Context) {
	var userDTO dtos.UserDTO

	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errors := dtos.Validate(userDTO)

	if len(errors) > 0 {
		ctx.JSON(http.StatusBadRequest, errors)
		return
	}

	handler.UseCase.Execute(userDTO)

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created"})
}
