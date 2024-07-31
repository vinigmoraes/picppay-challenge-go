package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"picpay-challenge-go/cmd/api/routes"
	"picpay-challenge-go/pkg/database"
	"picpay-challenge-go/pkg/dependency_injection"
)

func initRoutes(router *gin.Engine, database *gorm.DB) {
	routes.CreateUserRoute(dependency_injection.InjectCreateUserHandler(database), router)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	db := database.Init()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	initRoutes(r, db)

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
