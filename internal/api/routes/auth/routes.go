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
	router.GET("/users/info", controllers.UserController{}.UserInfo)
	router.GET("/users/inventory", controllers.UserController{}.UserInventory)
}

func post(router *gin.RouterGroup) {
	router.POST("/referrals/store-code", controllers.ReferralController{}.StoreCode)
}

func put(router *gin.RouterGroup) {
}
