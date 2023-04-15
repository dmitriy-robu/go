package services

import "go-rust-drop/internal/api/models"

type LevelService struct {
}

func (ls LevelService) GetLevelForByExperience(experience *int) models.Level {
	level := &models.Level{
		Current:       1,
		MinExperience: 0,
		MaxExperience: 100,
		Experience:    0,
	}

	if experience != nil {
		level.Experience = *experience
	}

	return *level
}
