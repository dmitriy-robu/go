package public

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/controllers"
)

func Routes(router *gin.RouterGroup) {
	get(router)
	post(router)
}

func get(router *gin.RouterGroup) {
	boxes := router.Group("/boxes")
	{
		boxes.GET("/", controllers.BoxController{}.Index)
		boxes.GET("/:uuid", controllers.BoxController{}.Show)
	}
  
	router.GET("/project-statistics", controllers.ProjectStatisticController{}.GetProjectStatistic)
	router.GET("/provably-fair", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
}

func post(router *gin.RouterGroup) {

}
