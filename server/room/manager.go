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
}

func CreateRoom(code string) *Room {
	mu.Lock()
	defer func() {
		mu.Unlock()
	}()

	nr := New(code)
	rooms[code] = nr
	return nr
}

func (r *Room) JoinUser(u *user.User) {
	r.mu.Lock()
	r.Users = append(r.Users, u)
	joinMsg := fmt.Sprint("[", u.Name, " joined the room]")
	fmt.Println(joinMsg)
	r.WriteMsgToUsers(fmt.Sprintf("%v - %v", date.NowDateAndTimeBR(), joinMsg))
	r.mu.Unlock()
}

func (r *Room) DisconnectUser(u *user.User) {
	r.mu.Lock()
	leftMsg := fmt.Sprint("[", u.Name, " left the room]")
	fmt.Println(leftMsg)
	idx := user.GetIndex(r.Users, u)
	r.Users = slices.Delete(r.Users, idx, idx+1)
	r.WriteMsgToUsers(fmt.Sprintf("%v - %v", date.NowDateAndTimeBR(), leftMsg))
	r.mu.Unlock()
	if len(r.Users) == 0 {
		delete(rooms, r.Code)
		r = nil
	}
}

func (r *Room) WriteMsgToUsers(m string) {
	for _, u := range r.Users {
		u.MsgBox <- m
	}
}
