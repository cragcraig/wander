package core

import "time"

type ActiveUser struct {
	Id        int
	Nick      string
	Conn      Connection
	timestamp time.Time
}

func (user ActiveUser) String() string {
	return user.Nick
}
