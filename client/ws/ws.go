package ws

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func CloseConn(c *websocket.Conn) {
	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("Error massaging close to ws (ws.CloseConn()):", err)
		log.Println("Forcing close ws (ws.CloseConn):", err)
		c.Close()
	}
	fmt.Println("Connection closed (ws.CloseConn())")
}

func ListenMsg(c *websocket.Conn) {
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Panic("Error to read message (ws.ListenMsg()):", err)
			return
		}
		fmt.Println(string(message))
	}
}
