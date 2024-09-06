package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picpay-challenge-go/cmd/api/dtos"
	errorhandler "picpay-challenge-go/cmd/api/error"
	"picpay-challenge-go/cmd/api/responses"
	usecase "picpay-challenge-go/internal/usecase/wallets"
)

type TransferMoneyHandler struct {
	UseCase usecase.TransferMoneyUseCase
}

func (handler *TransferMoneyHandler) TransferMoney(ctx *gin.Context) {
	var transferDTO dtos.TransferMoneyDTO

	if err := ctx.ShouldBind(&transferDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErrors := dtos.Validate(transferDTO)

	if len(validationErrors) > 0 {
		ctx.JSON(http.StatusBadRequest, validationErrors)
		return
	}

	wallet, err := handler.UseCase.Execute(transferDTO)

	if err != nil {
		message, statusCode := errorhandler.HandleError(err)
		ctx.JSON(statusCode, gin.H{"message": message})
		return
	}

	ctx.JSON(
		http.StatusCreated,
		responses.TransferMoneyResponse{UserID: wallet.UserID, Balance: wallet.Balance},
	)
}
