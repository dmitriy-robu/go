package test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-rust-drop/internal/api/controllers"
	"go-rust-drop/internal/api/resources"
	"go-rust-drop/internal/api/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProjectStatistic(t *testing.T) {
	gin.SetMode(gin.TestMode)

	projectStatisticManager := services.ProjectStatisticsManager{}

	r := gin.Default()
	controller := controllers.ProjectStatisticController{
		ProjectStatisticManager: projectStatisticManager,
	}
	r.GET("/project-statistics", controller.GetProjectStatistic)

	req, _ := http.NewRequest("GET", "/project-statistics", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.NoError(t, err)

	projectStatistic := projectStatisticManager.GetStatistics()
	expectedResponse := resources.ProjectStatisticResource{
		ProjectStatistic: projectStatistic,
	}
	expectedJSON, _ := expectedResponse.ToJSON()

	assert.Equal(t, expectedJSON, response)
}
