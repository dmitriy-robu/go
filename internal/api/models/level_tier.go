package models

type LevelTier struct {
	Name string `gorm:"column:name"`
	TierBox
}
