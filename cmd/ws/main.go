package main

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/ws"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/ws", func(c *gin.Context) {
		ws.Handler(c.Writer, c.Request)
	})

	err := router.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
