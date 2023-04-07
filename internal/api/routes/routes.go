package routes

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/controllers"
	"go-rust-drop/internal/api/middlewares"
	"go-rust-drop/internal/api/routes/auth"
	"go-rust-drop/internal/api/routes/public"
)

type Route interface {
}

func RouteHandle(router *gin.Engine) {
	router.GET("/auth/steam", controllers.SteamAuthController{}.Login)
	router.GET("/auth/steam/callback", controllers.SteamAuthController{}.Callback)

	publicGroup := router.Group("/api/v1")
	public.Routes(publicGroup)

	dashboardGroup := router.Group("/api/v1")
	steamMiddleware := middlewares.SteamMiddleware{}

	dashboardGroup.Use(steamMiddleware.VerifyToken)
	{
		auth.Routes(dashboardGroup)
	}
}
