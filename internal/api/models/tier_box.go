package models

import "time"

type TierBox struct {
	ID        uint   `gorm:"primaryKey"`
	UUID      string `gorm:"unique"`
	ImageURL  string `gorm:"column:image_url"`
	OpenTime  int    `gorm:"column:open_time,default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
