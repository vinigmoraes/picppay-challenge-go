package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	dbusers "picpay-challenge-go/pkg/database/users"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:postgres@localhost:5432/picpaychallenge"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.Table("users", &dbusers.UserModel{})

	return db
}
