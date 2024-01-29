package client

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"

	"github.com/gorilla/websocket"
)

var SERVER = "localhost:2345"
var READER = bufio.NewReader(os.Stdin)

func readInput() (string, error) {
	inputStr := ""
	inputStr, err := READER.ReadString('\n')
	if err != nil {
		log.Println("Error to read input (client.readInput()):", err)
		return "", err
	}

	inputStr = strings.TrimSuffix(inputStr, "\n")
	inputStr = strings.TrimSuffix(inputStr, "\r")
	return inputStr, nil
}

func readInputWithQuestion(question string) (string, error) {
	inputStr := ""
	var err error = nil

	for len(inputStr) == 0 || err != nil {
		if len(question) > 0 {
			fmt.Print(question)
		}

		inputStr, err = readInput()
		if err != nil {
			return "", err
		}
		if len(inputStr) == 0 {
			fmt.Println("")
			fmt.Println("Valor não pode ser vazio")
		}
	}

	return inputStr, err
}

func getInterruption() chan os.Signal {
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	return interrupt
}

func listenInterruption(cInterrupt chan os.Signal) {
	select {
	case <-cInterrupt:
		log.Panic("Caught interrupt signal (client.listenInterruption())")
		return
	}
}

func closeConn(c *websocket.Conn) {
	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("Error massaging close to ws (client.closeConn()):", err)
		log.Println("Forcing close ws (client.closeConn):", err)
		c.Close()
	}
	fmt.Println("Connection closed (client.closeConn())")
}

func listenWsMsg(c *websocket.Conn) {
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Panic("Error to read message (client.ReadMessage()):", err)
			return
		}
		fmt.Println(string(message))
	}
}

func manageClientMsg(c *websocket.Conn) {
	for {
		msg, err := readInputWithQuestion("")

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

	cInterrupt := getInterruption()
	defer func() {
		log.Println("Closing interruption (Exec.anon())")
		close(cInterrupt)
	}()

	userName, err := readInputWithQuestion("Digite seu nome: ")
	if err != nil {
		log.Panic("Error to read name (Exec.readInputWithQuestion()):", err)
	}

	room, err := readInputWithQuestion("Digite o código da sala: ")
	if err != nil {
		log.Panic("Error to read room (Exec.readInputWithQuestion()):", err)
	}
	go listenInterruption(cInterrupt)

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
	defer closeConn(c)

	fmt.Println("Connected")
	fmt.Println("-------------------------------")
	fmt.Println("")

	go listenWsMsg(c)
	manageClientMsg(c)
}
