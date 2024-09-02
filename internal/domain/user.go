package domain

import (
	"picpay-challenge-go/cmd/api/dtos"
	"time"
)

type User struct {
	ID                    int
	Name                  string
	Email                 string
	CPF                   string
	Password              string
	Status                string
	Type                  UserType
	WalletID              int
	CreatedAt             time.Time
	IsAbleToTransferMoney bool
}

func CreateUser(dto dtos.UserDTO) User {
	return User{
		Name:                  dto.Name,
		Email:                 dto.Email,
		CPF:                   dto.CPF,
		Password:              dto.Password,
		Status:                "active",
		CreatedAt:             time.Now(),
		Type:                  UserType(dto.Type),
		IsAbleToTransferMoney: setIsAbleToTransferMoney(dto.Type),
	}
}

func setIsAbleToTransferMoney(userType string) bool {
	if userType == COMUM.String() {
		return true
	}
	return false
}

func (u *User) setWalletId(walletId int) {
	if u.WalletID == 0 {
		u.WalletID = walletId
	}
}

func (u *User) SetId(id int) {
	if u.ID == 0 {
		u.ID = id
	}
}
