package auth

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/controllers"
)

func Routes(router *gin.RouterGroup) {
	get(router)
	post(router)
	put(router)
}

func get(router *gin.RouterGroup) {
	router.GET("/user/info", controllers.UserController{}.UserInfo)
}

func post(router *gin.RouterGroup) {

}

func put(router *gin.RouterGroup) {
}
