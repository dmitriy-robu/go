package models

import "gorm.io/gorm"

type UserBalance struct {
	gorm.Model
	ID      uint64 `gorm:"primaryKey"`
	UserID  uint64 `gorm:"unique"`
	Balance *int   `gorm:"column:balance"`
}
