package services

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
)

type LevelManager struct {
	levelRepository repositories.LevelRepository
}

func NewLevelManager(lr repositories.LevelRepository) LevelManager {
	return LevelManager{
		levelRepository: lr,
	}
}

func (ls LevelManager) GetLevelForByExperience(experience uint64) models.Level {
	var (
		err   error
		level models.Level
	)

	level, err = ls.levelRepository.GetLevelByExperience(experience)
	if err != nil {
		level.EndsAt = 1

		return level
	}

	return level
}
