package client

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/glauberratti/websoulcat/client/input"
	"github.com/glauberratti/websoulcat/client/interruption"
	"github.com/glauberratti/websoulcat/client/ws"
	"github.com/gorilla/websocket"
)

var SERVER = "localhost:2345"
var READER = bufio.NewReader(os.Stdin)

func manageClientMsg(c *websocket.Conn) {
	for {
		msg, err := input.ReadInputWithQuestion("")

		if err != nil {
			log.Panic("Error to read message (client.manageClientMsg()):", err)
			break
		}
		if len(msg) == 0 {
			continue
		}

		e := c.WriteMessage(websocket.TextMessage, []byte(msg))
		if e != nil {
			log.Panic("Error to write (client.manageClientMsg()):", err)
			return
		}
	}
}

func Exec() {
	fmt.Println("Executing client")
	fmt.Println("Connecting to ", SERVER)

	cInterrupt := interruption.GetInterruption()
	defer func() {
		log.Println("Closing interruption (Exec.anon())")
		close(cInterrupt)
	}()

	userName, err := input.ReadInputWithQuestion("Digite seu nome: ")
	if err != nil {
		log.Panic("Error to read name (Exec.ReadInputWithQuestion()):", err)
	}

	room, err := input.ReadInputWithQuestion("Digite o código da sala: ")
	if err != nil {
		log.Panic("Error to read room (Exec.ReadInputWithQuestion()):", err)
	}
	go interruption.ListenInterruption(cInterrupt)

	URL := url.URL{Scheme: "ws", Host: SERVER}
	q := URL.Query()
	q.Add("user", userName)
	q.Add("room", room)
	URL.RawQuery = q.Encode()

	fmt.Println("Connecting to WS")
	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Panic("Error to create dial (client.Exec()):", err)
		return
	}
	defer ws.CloseConn(c)

	fmt.Println("Connected")
	fmt.Println("-------------------------------")
	fmt.Println("")

	go ws.ListenMsg(c)
	manageClientMsg(c)
}
