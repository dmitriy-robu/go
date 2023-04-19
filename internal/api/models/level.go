package models

import "time"

type Level struct {
	ID          uint `gorm:"primaryKey"`
	Level       int  `gorm:"column:level"`
	StartsFrom  int  `gorm:"column:starts_from"`
	EndsAt      int  `gorm:"column:ends_at"`
	KeysGranted int  `gorm:"column:keys_granted"`
	LevelTierID int  `gorm:"column:level_tier_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type LevelData struct {
	Current       int
	MinExperience int
	MaxExperience int
	Experience    int
}
