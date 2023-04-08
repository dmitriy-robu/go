package models

import (
	"github.com/google/uuid"
	"go-rust-drop/internal/api/enum"
	"gorm.io/gorm"
)

type Item struct {
	ID              int                  `json:"id" gorm:"primaryKey"`
	UUID            uuid.UUID            `json:"uuid" gorm:"unique"`
	Name            string               `json:"name"`
	Price           float64              `json:"price"`
	Color           string               `json:"color"`
	GameEnvironment enum.GameEnvironment `json:"game_environment"`
	ImageUrl        string               `json:"image_url"`
	gorm.Model
}
