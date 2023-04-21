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
	user := router.Group("/users")
	{
		user.GET("/info", controllers.UserController{}.UserInfo)
		user.GET("/inventory", controllers.UserController{}.UserInventory)
		user.GET("/updatable-fields", controllers.UserController{}.GetUpdatableFields)
	}
	router.GET("/referrals/details", controllers.ReferralController{}.Details)
}

func post(router *gin.RouterGroup) {
	router.POST("/users/set-trade-url", controllers.UserController{}.StoreSteamTradeURL)
	router.POST("/referrals/store-code", controllers.ReferralController{}.StoreCode)
	router.POST("/boxes/open/:uuid", controllers.BoxController{}.Open)

	caseBattle := router.Group("/case-battles")
	{
		caseBattle.POST("/create", controllers.CaseBattleController{}.Create)
		caseBattle.POST("/:uuid/join", controllers.CaseBattleController{}.Join)
		caseBattle.POST("/:uuid/start", controllers.CaseBattleController{}.Start)
	}
}

func put(router *gin.RouterGroup) {
}
