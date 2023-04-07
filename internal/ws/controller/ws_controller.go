package controller

import (
	"context"
	"fmt"
	"go-rust-drop/internal/ws/services"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"time"
)

type WSController struct {
}

func (ws WSController) Ws(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println("Failed to accept ws connection:", err)
		return
	}
	defer func(c *websocket.Conn, code websocket.StatusCode, reason string) {
		err := c.Close(code, reason)
		if err != nil {
			log.Println("Failed to close ws connection:", err)
		}
	}(c, websocket.StatusNormalClosure, "")

	messages := make(chan string)

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	go services.WSService{}.HandleMessages(ctx, c, messages)

	for {
		select {
		case msg := <-messages:
			log.Println("Received message:", msg)
			messages <- fmt.Sprintf("You said: %s", msg)
		case <-time.After(time.Minute):
			log.Println("Closing connection due to inactivity")
			return
		}
	}
}
