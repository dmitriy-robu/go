package services

import "go-rust-drop/internal/api/models"

type LevelManager struct{}

func (lm LevelManager) GetLevelForByExperience(experience *int) models.Level {
	// Замените следующие строчки своей логикой для UserRepository
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
