package models

import "gorm.io/gorm"

type UserBalance struct {
	ID      uint64 `gorm:"primaryKey"`
	UserID  uint64 `gorm:"unique"`
	Balance *int
	gorm.Model
}

type UserBalanceWithUser struct {
	ID      uint64 `gorm:"primaryKey"`
	UserID  uint64 `gorm:"unique"`
	Balance *int
	User    User
	gorm.Model
}
