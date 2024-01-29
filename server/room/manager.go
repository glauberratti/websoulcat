package room

import (
	"fmt"
	"slices"
	"sync"

	"github.com/glauberratti/websoulcat/server/user"
	"github.com/glauberratti/websoulcat/utils/date"
)

var mu sync.Mutex

var rooms = make(map[string]*Room)

func GetRoom(code string) (*Room, bool) {
	mu.Lock()
	defer func() {
		mu.Unlock()
	}()
	r, ok := rooms[code]
	return r, ok

	// if ok {
	// 	return r
	// }
	// return createRoom(code)
}

func CreateRoom(code string) *Room {
	mu.Lock()
	defer func() {
		mu.Unlock()
	}()

	nr := New(code)
	// go nr.Run()
	rooms[code] = nr
	return nr
}

func (r *Room) JoinUser(u *user.User) {
	r.mu.Lock()
	r.Users = append(r.Users, u)
	joinMsg := fmt.Sprint("[", u.Name, " joined the room]")
	fmt.Println(joinMsg)
	// now := time.Now()
	// nowStr := fmt.Sprintf("%v:%v:%v", now.Hour(), now.Minute(), now.Second())
	for _, us := range r.Users {
		us.MsgBox <- fmt.Sprintf("%v - %v", date.NowDateAndTimeBR(), joinMsg)
	}
	r.mu.Unlock()
}

func (r *Room) DisconnectUser(u *user.User) {
	r.mu.Lock()
	leftMsg := fmt.Sprint("[", u.Name, " left the room]")
	fmt.Println(leftMsg)
	// now := time.Now()
	// nowStr := fmt.Sprintf("%v:%v:%v", now.Hour(), now.Minute(), now.Second())
	for _, us := range r.Users {
		us.MsgBox <- fmt.Sprintf("%v - %v", date.NowDateAndTimeBR(), leftMsg)
	}
	idx := user.GetIndex(r.Users, u)
	// r.Users = append(r.Users[:idx], r.Users[idx+1:]...)
	slices.Delete(r.Users, idx, idx+1)
	r.mu.Unlock()
}

// func readMsgFromUser(r *Room) {

// }

func (r *Room) WriteMsgToUsers(m string) {
	for _, u := range r.Users {
		u.MsgBox <- m
	}
}
