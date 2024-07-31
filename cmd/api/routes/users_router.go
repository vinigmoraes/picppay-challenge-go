package routes

import (
	"github.com/gin-gonic/gin"
	"picpay-challenge-go/cmd/api/handlers"
)

func CreateUserRoute(handler handlers.CreateUserHandler, router *gin.Engine) {
	router.POST("/users", handler.CreateUser)
}
