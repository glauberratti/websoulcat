package user

import (
	"fmt"
	"sync"
)

var id int
var once sync.Once

func init() {
	once.Do(func() {
		id = 0
	})
}

func nextId() int {
	id++
	return id
}

type User struct {
	ID     int
	Name   string
	MsgBox chan string
}

func New(name string) *User {
	id := nextId()
	user := &User{
		ID:     id,
		Name:   name,
		MsgBox: make(chan string, 5),
	}

	if len(user.Name) == 0 {
		name = fmt.Sprint("User", id)
	}

	return user
}

func GetIndex(us []*User, user *User) int {
	for i, u := range us {
		if u.ID == user.ID {
			return i
		}
	}
	return -1
}
