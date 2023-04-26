package services

import (
	"go-rust-drop/config"
	"go-rust-drop/internal/api/models"
	"math/rand"
	"strconv"
	"time"
)

type ProjectStatisticsManager struct {
}

func NewProjectStatisticsManager() ProjectStatisticsManager {
	return ProjectStatisticsManager{}
}

const (
	DividerForOpenedCaseTime  = 10000
	DividerForOnlineUsersTime = 10000000
)

func (psm ProjectStatisticsManager) GetStatistics() models.ProjectStatistic {
	var projectStatistic models.ProjectStatistic

	projectStatisticConfig := config.SetProjectStatistic()

	projectStatistic = models.ProjectStatistic{
		CasesOpened: psm.getOpenedCases(projectStatisticConfig.OpenedCasesStartValue),
		TotalUsers:  psm.getTotalUsers(projectStatisticConfig.TotalUsersStartValue),
		OnlineUsers: psm.getOnlineUsers(projectStatisticConfig.OnlineUsersStartValue),
	}

	return projectStatistic
}

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(1))
}

func (psm ProjectStatisticsManager) getOpenedCases(startValueOpenedCases int) int {
	timeNow := time.Now().Unix()
	lastThreeNumbersFromTimestamp, _ := strconv.Atoi(strconv.FormatInt(timeNow, 10)[len(strconv.FormatInt(timeNow, 10))-3:])
	dividedCurrentTime := int(float64(timeNow) / DividerForOpenedCaseTime)

	return startValueOpenedCases + lastThreeNumbersFromTimestamp + dividedCurrentTime
}

func (psm ProjectStatisticsManager) getTotalUsers(startValueTotalUsers int) int {
	timeNow := strconv.FormatInt(time.Now().Unix(), 10)
	oneNumberBeforeLatestFromTime, _ := strconv.Atoi(timeNow[len(timeNow)-2 : len(timeNow)-1])
	twoNumbersAfterThreeFirstNumbers, _ := strconv.Atoi(timeNow[3:5])

	return startValueTotalUsers + oneNumberBeforeLatestFromTime + twoNumbersAfterThreeFirstNumbers
}

func (psm ProjectStatisticsManager) getOnlineUsers(startValueOnlineUsers int) int {
	dividedCurrentTime := int(float64(time.Now().Unix()) / DividerForOnlineUsersTime)
	tenthPartFromRandomNumber := newRand().Float64()*3 + 1

	return startValueOnlineUsers + int(float64(dividedCurrentTime)*tenthPartFromRandomNumber)
}
