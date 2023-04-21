package repositories

import (
	"go-rust-drop/internal/api/models"
	"gorm.io/gorm"
)

type LevelRepository struct {
	MysqlDB *gorm.DB
}

func (lr LevelRepository) GetLevelByExperience(experience int) (models.Level, error) {
	var (
		err   error
		level models.Level
	)

	err = lr.MysqlDB.Where("starts_from <= ?", experience).Where("ends_at >= ?", experience).First(&level).Error
	if err != nil {
		return level, err
	}

	return level, nil
}
