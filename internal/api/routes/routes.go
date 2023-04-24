package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/controllers"
	"go-rust-drop/internal/api/database/mongodb"
	"go-rust-drop/internal/api/middlewares"
	"go-rust-drop/internal/api/routes/auth"
	"go-rust-drop/internal/api/routes/public"
	"log"
)

type Route interface {
}

func RouteHandle(router *gin.Engine, controllersInstance controllers.Controllers) {
	store, err := mongodb.InitMongoSessionStore()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB session store: %v", err)
	}
	router.Use(sessions.Sessions("session_name", store))

	authSteam := router.Group("/auth/steam")
	{
		authSteam.GET("/login", controllersInstance.SteamAuthController.Login)
		authSteam.GET("/callback", controllersInstance.SteamAuthController.Callback)
	}

	publicGroup := router.Group("/api/v1")
	public.Routes(publicGroup)

	authGroup := router.Group("/api/v1")
	Middleware := middlewares.Middleware{}

	authGroup.Use(Middleware.AuthRequired)
	{
		auth.Routes(authGroup, controllersInstance)
	}
}
