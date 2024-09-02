package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	dbusers "picpay-challenge-go/pkg/database/users"
	dbwallets "picpay-challenge-go/pkg/database/wallet"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:postgres@localhost:5432/picpaychallenge"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(dbusers.Users{}, dbwallets.Wallets{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
