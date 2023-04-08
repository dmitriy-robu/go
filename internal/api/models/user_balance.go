package models

var TableUserBalance = "user_balance"

type UserBalance struct {
	ID      uint64 `gorm:"primaryKey"`
	UserID  uint64 `gorm:"unique"`
	Balance *int
	User    User
}
