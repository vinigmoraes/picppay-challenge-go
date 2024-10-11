package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
	dbusers "picpay-challenge-go/pkg/database/users"
	dbwallets "picpay-challenge-go/pkg/database/wallet"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:postgres@postgres:5432/picpaychallenge?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		slog.Warn("error", err)
		return Init()
	}

	err = db.AutoMigrate(dbusers.Users{}, dbwallets.Wallets{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
