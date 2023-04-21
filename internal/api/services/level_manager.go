package services

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
)

type LevelManager struct {
	levelRepository repositories.LevelRepository
}

func (ls LevelManager) GetLevelForByExperience(experience int) models.Level {
	var (
		err   error
		level models.Level
	)

	ls.levelRepository.MysqlDB = MysqlDB

	level, err = ls.levelRepository.GetLevelByExperience(experience)
	if err != nil {
		level.EndsAt = 1

		return level
	}

	return level
}
