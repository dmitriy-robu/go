package services

import (
	"context"
	"fmt"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type WSService struct {
}

func (wss WSService) ReadMessages(ctx context.Context, c *websocket.Conn, incoming chan string) {
	for {
		var message string
		err := wsjson.Read(ctx, c, &message)
		if err != nil {
			fmt.Println("Failed to read ws message:", err)
			break
		}

		incoming <- message
	}
}

func (wss WSService) WriteMessages(ctx context.Context, c *websocket.Conn, outgoing chan string) {
	for {
		msg := <-outgoing
		err := wsjson.Write(ctx, c, msg)
		if err != nil {
			fmt.Println("Failed to write ws message:", err)
			break
		}
	}
}
