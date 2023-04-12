package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/steam"
	"go-rust-drop/internal/api/database/migrations"
	"go-rust-drop/internal/api/database/mongodb"
	"go-rust-drop/internal/api/routes"
	"log"
	"os"
)

func main() {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}
	r := gin.Default()

	goth.UseProviders(
		steam.New(os.Getenv("STEAM_KEY"), os.Getenv("STEAM_CALLBACK_URL")),
	)

	mongoStore, err := mongodb.InitMongoSessionStore()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB session store: %v", err)
	}

	r.Use(sessions.Sessions("mysession", mongoStore))

	routes.RouteHandle(r)

	go migrations.Migrations{}.MigrateAll()

	if err = r.Run(":" + os.Getenv("GO_PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}
