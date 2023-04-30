package auth

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/controllers"
)

func Routes(router *gin.RouterGroup, controllersInstance controllers.Controllers) {
	get(router, controllersInstance)
	post(router, controllersInstance)
	put(router, controllersInstance)

}

func get(router *gin.RouterGroup, controllersInstance controllers.Controllers) {
	user := router.Group("/users")
	{
		user.GET("/info", controllersInstance.UserController.UserInfo)
		user.GET("/inventory", controllersInstance.UserController.UserInventory)
		user.GET("/updatable-fields", controllersInstance.UserController.GetUpdatableFields)
	}

	router.GET("/referrals/details", controllersInstance.ReferralController.Details)
}

func post(router *gin.RouterGroup, controllersInstance controllers.Controllers) {
	router.POST("/users/set-trade-url", controllersInstance.UserController.StoreSteamTradeURL)
	router.POST("/referrals/store-code", controllersInstance.ReferralController.StoreCode)

	caseBattle := router.Group("/case-battles")
	{
		caseBattle.POST("/create", controllersInstance.CaseBattleController.Create)
		caseBattle.POST("/:uuid/join", controllersInstance.CaseBattleController.Join)
		caseBattle.POST("/:uuid/start", controllersInstance.CaseBattleController.Start)
	}
}

func put(router *gin.RouterGroup, controllersInstance controllers.Controllers) {
}
