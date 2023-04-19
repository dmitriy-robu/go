package repositories

import "go-rust-drop/internal/api/models"

type LevelRepository struct {
}

func (lr LevelRepository) GetLevelByExperience(experience int) (models.Level, error) {
	var (
		err   error
		level models.Level
	)

	err = MysqlDB.Where("starts_from <= ?", experience).Where("ends_at >= ?", experience).First(&level).Error
	if err != nil {
		return level, err
	}

	return level, nil
}
