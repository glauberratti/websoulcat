package chat

import (
	"fmt"
	"net/http"

	"github.com/glauberratti/websoulcat/server/room"
	pkgRoom "github.com/glauberratti/websoulcat/server/room"
	"github.com/glauberratti/websoulcat/server/user"
	"github.com/glauberratti/websoulcat/utils/date"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	roomCode := r.URL.Query().Get("room")
	userName := r.URL.Query().Get("user")
	// fmt.Println("Passou nos query params")

	room, roomExists := pkgRoom.GetRoom(roomCode)
	if !roomExists {
		room = pkgRoom.CreateRoom(roomCode)
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}
	// fmt.Println("Connected")
	// fmt.Println("-------------------------------")
	// fmt.Println("")

	// N(userName, room, conn)

	u := user.New(userName)
	// room.Join <- u
	room.JoinUser(u)

	go writeMsgToClients(u, conn)
	go readClientMsg(u, conn, room)
}

func readClientMsg(u *user.User, conn *websocket.Conn, r *room.Room) {
	defer func() {
		// r.Leave <- u
		r.DisconnectUser(u)
		conn.Close()
	}()

	for {
		_, msgB, err := conn.ReadMessage()
		if err != nil {
			break
		}
		msg := string(msgB)
		// now := time.Now()
		msg = fmt.Sprintf("%v - %v: %v", date.NowDateAndTimeBR(), u.Name, msg)
		// nowStr := fmt.Sprintf("%v:%v:%v", now.Hour(), now.Minute(), now.Second())
		// r.MsgBox <- msg
		// r.DisconnectUser(u)
		r.WriteMsgToUsers(msg)
	}
}

func writeMsgToClients(u *user.User, conn *websocket.Conn) {
	defer conn.Close()
	for msg := range u.MsgBox {
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			break
		}
	}
}

// func N(o ...any) {

// }
