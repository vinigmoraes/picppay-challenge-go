package routes

import (
	"github.com/gin-gonic/gin"
	"picpay-challenge-go/cmd/api/handlers"
)

func TransferMoney(handler handlers.TransferMoneyHandler, router *gin.Engine) {
	router.POST("/transfer", handler.TransferMoney)
}
