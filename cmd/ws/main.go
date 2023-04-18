package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-rust-drop/internal/ws/controller"
	"log"
	"os"
)

func main() {
	var err error

	if err = godotenv.Load(os.Getenv("ROOT_PATH") + "/.env"); err != nil {

		log.Fatalln(err)
		return
	}

	router := gin.Default()
	router.GET("/ws", func(c *gin.Context) {
		controller.WSController{}.Ws(c.Writer, c.Request)
	})

	err = router.Run(os.Getenv("WS_PORT"))
	if err != nil {
		log.Fatalln(err)
		return
	}
}
