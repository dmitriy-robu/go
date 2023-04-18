package models

import (
	"go-rust-drop/internal/api/enum"
	"gorm.io/gorm"
)

type BoxItem struct {
	ID     uint               `json:"ID" gorm:"primaryKey"`
	BoxID  uint               `json:"box_id" gorm:"foreignId"`
	ItemID uint               `json:"item_id" gorm:"foreignId"`
	Rarity enum.BoxItemRarity `json:"rarity"`
	gorm.Model
}
