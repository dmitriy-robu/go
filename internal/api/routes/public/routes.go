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
}

func post(router *gin.RouterGroup) {

}
