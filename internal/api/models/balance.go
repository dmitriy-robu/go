package models

import "gorm.io/gorm"

type UserBalance struct {
	gorm.Model
	UserID  uint `gorm:"unique"`
	Balance int  `gorm:"column:balance"`
}
