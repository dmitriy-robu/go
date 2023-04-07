package services

import (
	"context"
	"log"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type WSService struct {
}

func (wss WSService) HandleMessages(ctx context.Context, c *websocket.Conn, messages chan string) {
	for {
		var message string
		err := wsjson.Read(ctx, c, &message)
		if err != nil {
			log.Println("Failed to read ws message:", err)
			break
		}

		messages <- message

		response := <-messages
		err = wsjson.Write(ctx, c, response)
		if err != nil {
			log.Println("Failed to write ws message:", err)
			break
		}
	}
}
