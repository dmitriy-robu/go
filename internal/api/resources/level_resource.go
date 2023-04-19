package resources

import (
	"go-rust-drop/internal/api/models"
)

type LevelResource struct {
	Level models.Level
}

func (l *LevelResource) ToJSON() (map[string]interface{}, error) {
	var level map[string]interface{}

	level = map[string]interface{}{
		"current":        l.Level.Level,
		"min_experience": l.Level.StartsFrom,
		"max_experience": l.Level.EndsAt,
	}

	return level, nil
}
