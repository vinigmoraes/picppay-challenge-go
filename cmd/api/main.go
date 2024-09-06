package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
	"net/http"
	"picpay-challenge-go/cmd/api/routes"
	"picpay-challenge-go/pkg/database"
	"picpay-challenge-go/pkg/dependency_injection"
	"picpay-challenge-go/pkg/messagequeue"
)

func initRoutes(router *gin.Engine, database *gorm.DB, amqp *amqp091.Channel) {
	routes.CreateUserRoute(dependency_injection.InjectCreateUserHandler(database), router)
	routes.TransferMoney(dependency_injection.InjectTransferMoneyHandler(database, amqp), router)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	db := database.Init()
	broker, _ := messagequeue.Init()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	initRoutes(r, db, broker)

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
