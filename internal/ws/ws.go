package ws

import (
	"context"
	"fmt"
	"net/http"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		fmt.Println("Failed to accept ws connection:", err)
		return
	}
	defer func(c *websocket.Conn, code websocket.StatusCode, reason string) {
		err := c.Close(code, reason)
		if err != nil {
			fmt.Println("Failed to close ws connection:", err)
		}
	}(c, websocket.StatusNormalClosure, "")

	ctx, cancel := context.WithTimeout(r.Context(), time.Minute)
	defer cancel()

	for {
		var message string
		err := wsjson.Read(ctx, c, &message)
		if err != nil {
			fmt.Println("Failed to read ws message:", err)
			break
		}

		fmt.Println("Received message:", message)

		err = wsjson.Write(ctx, c, fmt.Sprintf("You said: %s", message))
		if err != nil {
			fmt.Println("Failed to write ws message:", err)
			break
		}
	}
}
