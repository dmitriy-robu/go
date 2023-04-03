package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-rustdrop/m/internal/routes"
	"log"
	"os"
)

func main() {
	var err error

	if err = godotenv.Load(os.Getenv("ROOT_PATH") + "/.env"); err != nil {
		log.Fatalln(err)
		return
	}

	r := gin.Default()
	routes.SetRoutes(r)

	if err = r.Run(os.Getenv("GO_PORT")); err != nil {
		log.Fatalln(err)
		return
	}
}
