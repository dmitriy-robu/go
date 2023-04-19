package models

import "time"

type TierKey struct {
	ID        uint      `gorm:"primaryKey"`
	LevelTier LevelTier `gorm:"foreignKey:LevelTierID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
