package resources

import (
	"go-rust-drop/internal/api/models"
)

type ProjectStatisticResource struct {
	ProjectStatistic models.ProjectStatistic
}

func (psr *ProjectStatisticResource) ToJSON() (map[string]interface{}, error) {
	var level map[string]interface{}

	level = map[string]interface{}{
		"cases_opened": psr.ProjectStatistic.CasesOpened,
		"total_users":  psr.ProjectStatistic.TotalUsers,
		"online_users": psr.ProjectStatistic.OnlineUsers,
	}

	return level, nil
}
