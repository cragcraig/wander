package core

import (
	"strconv"
	"time"
)

type ActiveUser struct {
	Id        int
    Nick      string
	Conn      Connection
	timestamp time.Time
}

func AuthNewUsers(c <-chan Connection, users chan<- ActiveUser) {
	defer close(users)
    nextId := 0
	for conn := range c {
		go authConnection(nextId, conn, users)
        nextId++
	}
}

func authConnection(id int, conn Connection, users chan<- ActiveUser) {
	conn.Write <- "Welcome to a Science Fiction Universe!\n"
	for {
		conn.Write <- "By what name would you like to be known?\n"
		nick, ok := conn.Prompt()
		if !ok {
			// Client disconnected.
			return
		}
        nick = strings.TrimSpace(nick)
		if nick != "" {
			users <- ActiveUser{id, nick, conn, time.Now()}
			return
		}
		conn.Write <- "Bad answer...\n"
	}
}
