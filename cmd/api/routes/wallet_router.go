package routes

import (
	"github.com/gin-gonic/gin"
	"picpay-challenge-go/cmd/api/consumers"
	"picpay-challenge-go/cmd/api/handlers"
)

func TransferMoney(handler handlers.TransferMoneyHandler, router *gin.Engine) {
	router.POST("/transfer", handler.TransferMoney)
}

func PutMoneyInWallet(handler consumers.TransferMoneyConsumerHandler) {
	handler.Consume()
}

func GiveMoneyBackToPayer(consumer consumers.GiveMoneyBackConsumer) {
	consumer.Consume()
}
