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
	router.GET("/users/updatable-fields", controllers.UserController{}.GetUpdatableFields)
	router.GET("/referrals/details", controllers.ReferralController{}.Details)
}

func post(router *gin.RouterGroup) {
	router.POST("/users/set-trade-url", controllers.UserController{}.StoreSteamTradeURL)
	router.POST("/referrals/store-code", controllers.ReferralController{}.StoreCode)
}

func put(router *gin.RouterGroup) {
}
