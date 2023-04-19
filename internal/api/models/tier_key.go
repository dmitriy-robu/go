package models

import "time"

type TierKey struct {
	ID        uint      `gorm:"primaryKey"`
	LevelTier LevelTier `gorm:"foreignKey:LevelTierID"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}
