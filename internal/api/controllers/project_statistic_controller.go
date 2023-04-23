package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/resources"
	"go-rust-drop/internal/api/services"
	"go-rust-drop/internal/api/utils"
	"net/http"
)

type ProjectStatisticController struct {
	errorHandler            utils.Errors
	ProjectStatisticManager services.ProjectStatisticsManager
}

func (psc ProjectStatisticController) GetProjectStatistic(c *gin.Context) {
	var (
		projectStatistic         models.ProjectStatistic
		projectStatisticResource resources.ProjectStatisticResource
	)

	projectStatistic = psc.ProjectStatisticManager.GetStatistics()

	projectStatisticResource = resources.ProjectStatisticResource{
		ProjectStatistic: projectStatistic,
	}

	projectStatisticJSON, err := projectStatisticResource.ToJSON()
	if err != nil {
		psc.errorHandler.HandleError(c, http.StatusInternalServerError, "Error converting project statistic to JSON", err)
		return
	}

	c.JSON(http.StatusOK, projectStatisticJSON)
}
