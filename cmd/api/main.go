package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-rust-drop/internal/api/controllers"
	"go-rust-drop/internal/api/database/migrations"
	"go-rust-drop/internal/api/database/mysql"
	"go-rust-drop/internal/api/routes"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	var err error

	if err = godotenv.Load(os.Getenv("ROOT_PATH") + "/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}

	container := NewContainer()

	err = container.Provide(func() *gorm.DB { return MysqlDB })
	//err = container.Provide(func() *mongo.Database { return mongodb })

	var controllersInstance controllers.Controllers
	err = container.Invoke(func(c controllers.Controllers) {
		controllersInstance = c
	})

	if err != nil {
		log.Fatalf("Failed to invoke controllers: %v", err)
	}

	r := gin.Default()

	routes.RouteHandle(r, controllersInstance)

	go migrations.Migrations{}.MigrateAll()

	if err = r.Run(":" + os.Getenv("GO_PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}

var MysqlDB = getMysqlDb()

func getMysqlDb() *gorm.DB {
	db, err := mysql.GetGormConnection()
	if err != nil {
		log.Fatalln("Error getting MySQL connection")
	}

	return db
}
