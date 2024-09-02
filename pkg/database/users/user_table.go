package dbusers

import "time"

type Users struct {
	ID                    int `gorm:"primary_key"`
	Name                  string
	Email                 string `gorm:"unique"`
	CPF                   string `gorm:"unique"`
	Password              string
	Status                string `gorm:"default:'active'"`
	Type                  string
	CreatedAt             time.Time
	IsAbleToTransferMoney bool
}
