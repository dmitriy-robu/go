package models

import (
	"go-rust-drop/internal/api/enum"
)

type BoxItems []BoxItem

type BoxItem struct {
	BoxID  uint               `json:"boxId" gorm:"primaryKey"`
	ItemID uint               `json:"itemId" gorm:"primaryKey"`
	Rarity enum.BoxItemRarity `json:"rarity"`
}