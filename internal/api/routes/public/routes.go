package public

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/controllers"
)

func Routes(router *gin.RouterGroup, controllersInstance controllers.Controllers) {
	get(router, controllersInstance)
	post(router, controllersInstance)
}

func get(router *gin.RouterGroup, controllersInstance controllers.Controllers) {
	boxes := router.Group("/boxes")
	{
		boxes.GET("/", controllersInstance.BoxController.Index)
		boxes.GET("/:uuid", controllersInstance.BoxController.Show)
	}

	router.GET("/project-statistics", controllersInstance.ProjectStatisticController.GetProjectStatistic)
	router.GET("/provably-fair", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
}

func post(router *gin.RouterGroup, controllersInstance controllers.Controllers) {
	openBox := router.Group("/boxes")
	{
		openBox.POST("/:uuid/open", controllersInstance.OpenBoxController.Open)
	}
}
