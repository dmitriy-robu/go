package config

import "go-rust-drop/internal/api/utils"

type ProjectStatistic struct {
	OpenedCasesStartValue int
	TotalUsersStartValue  int
	OnlineUsersStartValue int
}

func SetProjectStatistic() ProjectStatistic {
	env := utils.Environment{}
	return ProjectStatistic{
		OpenedCasesStartValue: env.GetEnvOrDefault("STATISTICS_OPENED_CASES_START_VALUE", 3000000).(int),
		TotalUsersStartValue:  env.GetEnvOrDefault("STATISTICS_TOTAL_USERS_START_VALUE", 2000).(int),
		OnlineUsersStartValue: env.GetEnvOrDefault("STATISTICS_ONLINE_USERS_START_VALUE", 300).(int),
	}
}
