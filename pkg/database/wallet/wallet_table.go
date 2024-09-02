package dbwallets

type Wallets struct {
	ID      int     `gorm:"primary_key"`
	UserID  int     `gorm:"foreignkey:UserID"`
	Balance float64 `gorm:"not null"`
}
