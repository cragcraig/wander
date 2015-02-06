package core

import "time"

// Must be safe to copy-by-value
type ActiveUser struct {
	Id        int
	Nick      string
	Conn      *Connection
	timestamp time.Time
}

func (user *ActiveUser) String() string {
	return user.Nick
}
