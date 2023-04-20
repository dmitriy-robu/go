package models

import (
	"github.com/google/uuid"
	"go-rust-drop/internal/api/enum"
	"gorm.io/gorm"
)

type Items []Item

type Item struct {
	ID              uint                 `json:"id" gorm:"primaryKey" faker:"-"`
	UUID            uuid.UUID            `json:"uuid" gorm:"unique"`
	Name            string               `json:"name" gorm:"column:name"`
	Price           int                  `json:"price" gorm:"column:price"`
	Color           string               `json:"color" gorm:"column:color"`
	GameEnvironment enum.GameEnvironment `json:"game_environment"`
	ImageUrl        string               `json:"image_url"`
	gorm.Model
}
