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
	ProjectStatisticManager services.ProjectStatisticsManager
}

func NewProjectStatisticController(
	projectStatisticManager services.ProjectStatisticsManager,
) ProjectStatisticController {
	return ProjectStatisticController{
		ProjectStatisticManager: projectStatisticManager,
	}
}

func (psc ProjectStatisticController) GetProjectStatistic(c *gin.Context) {
	var (
		projectStatistic         models.ProjectStatistic
		projectStatisticResource resources.ProjectStatisticResource
		errorHandler             utils.Errors
	)

	projectStatistic = psc.ProjectStatisticManager.GetStatistics()

	projectStatisticResource = resources.ProjectStatisticResource{
		ProjectStatistic: projectStatistic,
	}

	projectStatisticJSON, err := projectStatisticResource.ToJSON()
	if err != nil {
		errorHandler = utils.Errors{
			Code:    http.StatusInternalServerError,
			Message: "Error converting project statistic to JSON",
			Err:     err,
		}
		errorHandler.HandleError(c)
		return
	}

	c.JSON(http.StatusOK, projectStatisticJSON)
}
