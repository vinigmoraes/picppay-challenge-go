package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"os"
	"picpay-challenge-go/cmd/api/metrics"
	"picpay-challenge-go/cmd/api/routes"
	"picpay-challenge-go/pkg/database"
	"picpay-challenge-go/pkg/dependency_injection"
	"picpay-challenge-go/pkg/messagequeue"
)

func initRoutes(router *gin.Engine, database *gorm.DB, amqp *amqp091.Channel) {
	routes.CreateUserRoute(dependency_injection.InjectCreateUserHandler(database), router)
	routes.TransferMoney(dependency_injection.InjectTransferMoneyHandler(database, amqp), router)
}

func initConsumers(database *gorm.DB, amqp *amqp091.Channel) {
	routes.PutMoneyInWallet(dependency_injection.InjectTransferMoneyConsumer(database, amqp))
	routes.GiveMoneyBackToPayer(dependency_injection.InjectGiveMoneyBackConsumer(database, amqp))
}

func setLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	db := database.Init()

	broker, _ := messagequeue.Init()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	initRoutes(r, db, broker)
	initConsumers(db, broker)
	setLogger()

	return r
}

func monitor() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		_ = http.ListenAndServe(":2112", nil)
	}()
}

func initMetrics() {
	metrics.InitTransactionsMetrics()
}

func main() {
	r := setupRouter()
	monitor()
	initMetrics()

	r.Run(":8080")
}
