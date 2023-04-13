package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-rust-drop/internal/api/database/migrations"
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

	routes.RouteHandle(r)

	go migrations.Migrations{}.MigrateAll()

	if err = r.Run(":" + os.Getenv("GO_PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}
