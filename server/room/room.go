package room

import (
	"sync"

	"github.com/glauberratti/websoulcat/server/user"
)

type Room struct {
	Users []*user.User
	Code  string
	mu    sync.Mutex
}

func New(code string) *Room {
	return &Room{
		Users: make([]*user.User, 0),
		Code:  code,
	}
}
