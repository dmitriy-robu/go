package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/controllers"
	"go-rust-drop/internal/api/middlewares"
	"go-rust-drop/internal/api/routes/auth"
	"go-rust-drop/internal/api/routes/public"
	"os"
)

type Route interface {
}

func RouteHandle(router *gin.Engine) {
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	router.Use(sessions.Sessions("session_name", store))

	router.GET("/auth/steam", controllers.SteamAuthController{}.Login)
	router.GET("/auth/steam/callback", controllers.SteamAuthController{}.Callback)

	publicGroup := router.Group("/api/v1")
	public.Routes(publicGroup)

	authGroup := router.Group("/api/v1")
	steamMiddleware := middlewares.SteamMiddleware{}

	authGroup.Use(steamMiddleware.AuthRequired)
	{
		auth.Routes(authGroup)
	}
}
