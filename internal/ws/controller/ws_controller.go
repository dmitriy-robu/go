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

	incoming := make(chan string)
	outgoing := make(chan string)

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	go services.WSService{}.ReadMessages(ctx, c, incoming)
	go services.WSService{}.WriteMessages(ctx, c, outgoing)

	for {
		select {
		case msg := <-incoming:
			log.Println("Received message:", msg)
			outgoing <- fmt.Sprintf("You said: %s", msg)
		case <-time.After(time.Minute):
			log.Println("Closing connection due to inactivity")
			return
		}
	}
}
