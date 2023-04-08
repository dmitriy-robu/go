package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/steam"
	"go-rust-drop/internal/api/routes"
	"log"
	"os"
)

func main() {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
		return
	}
	/*
		if err = godotenv.Load(os.Getenv("ROOT_PATH") + "/.env"); err != nil {
			log.Fatalln(err)
			return
		}

	*/
	r := gin.Default()

	goth.UseProviders(
		steam.New(os.Getenv("STEAM_KEY"), os.Getenv("STEAM_CALLBACK_URL")),
	)

	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))

	r.Use(sessions.Sessions("mysession", store))

	routes.RouteHandle(r)

	if err = r.Run(":" + os.Getenv("GO_PORT")); err != nil {
		log.Fatalln(err)
		return
	}
}
