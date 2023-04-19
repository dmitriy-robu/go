package services

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
)

type LevelService struct {
	levelRepository repositories.LevelRepository
}

func (ls LevelService) GetLevelForByExperience(experience int) models.Level {
	var (
		err   error
		level models.Level
	)

	level, err = ls.levelRepository.GetLevelByExperience(experience)
	if err != nil {
		level.Level = 0
		level.StartsFrom = 0
		level.EndsAt = 1

		return level
	}

	return level
}
