package room

import (
	"sync"

	"github.com/glauberratti/websoulcat/server/user"
)

// type Message struct {
// 	Author string
// 	Msg    string
// }

type Room struct {
	Users []*user.User
	Code  string
	// Join   chan *user.User
	// Leave  chan *user.User
	// MsgBox chan string
	mu sync.Mutex
}

func New(code string) *Room {
	return &Room{
		Users: make([]*user.User, 0),
		Code:  code,
		// Join:   make(chan *user.User),
		// Leave:  make(chan *user.User),
		// MsgBox: make(chan string),
	}
}

// func (r *Room) Run() {
// 	for {
// 		select {
// 		case us := <-r.Join:
// 			r.mu.Lock()
// 			fmt.Println(us.Name, "joined the room")
// 			now := time.Now()
// 			nowStr := fmt.Sprintf("%v:%v:%v", now.Hour(), now.Minute(), now.Second())
// 			for _, u := range r.Users {
// 				u.MsgBox <- fmt.Sprintf("%v - %v joined the room", nowStr, us.Name)
// 			}
// 			r.Users = append(r.Users, us)
// 			r.mu.Unlock()

// 		case us := <-r.Leave:
// 			r.mu.Lock()
// 			fmt.Println(us.Name, "left the room")
// 			now := time.Now()
// 			nowStr := fmt.Sprintf("%v:%v:%v", now.Hour(), now.Minute(), now.Second())
// 			for _, u := range r.Users {
// 				u.MsgBox <- fmt.Sprintf("%v - %v left the room", nowStr, us.Name)
// 			}
// 			idx := user.GetIndex(r.Users, us)
// 			// r.Users = append(r.Users[:idx], r.Users[idx+1:]...)
// 			slices.Delete(r.Users, idx, idx+1)
// 			r.mu.Unlock()

// 		case msg := <-r.MsgBox:
// 			for _, u := range r.Users {
// 				u.MsgBox <- msg
// 			}
// 		}
// 	}
// }
