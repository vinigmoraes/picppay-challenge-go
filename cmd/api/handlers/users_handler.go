package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picpay-challenge-go/cmd/api/dtos"
	errorhandler "picpay-challenge-go/cmd/api/error"
	usecase "picpay-challenge-go/internal/usecase/users"
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

	validationErrors := dtos.Validate(userDTO)

	if len(validationErrors) > 0 {
		ctx.JSON(http.StatusBadRequest, validationErrors)
		return
	}

	user, err := handler.UseCase.Execute(userDTO)

	if err != nil {
		message, statusCode := errorhandler.HandleError(err)
		ctx.JSON(statusCode, gin.H{"message": message})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user_id": user.ID})
}
