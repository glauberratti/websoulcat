package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/glauberratti/websoulcat/server/chat"
	"github.com/glauberratti/websoulcat/utils/date"
)

var SERVER = "localhost:2345"

func Exec() {
	fmt.Println("Executing server")
	http.HandleFunc("/", chat.HandleConnection)
	fmt.Println("Server started at", SERVER)
	fmt.Println(date.NowDateAndTimeBR())
	fmt.Println("-------------------------------")
	fmt.Println("")
	err := http.ListenAndServe(SERVER, nil)
	if err != nil {
		log.Fatal(err)
	}
}
