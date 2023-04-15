package models

import "gorm.io/gorm"

type UserBalance struct {
	gorm.Model
	UserID  uint64 `gorm:"unique"`
	Balance *int   `gorm:"column:balance"`
}
