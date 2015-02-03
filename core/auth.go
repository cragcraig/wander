package core

import (
	"fmt"
	"strings"
	"time"
)

func AuthNewUsers(c <-chan Connection, users chan<- ActiveUser) {
	defer close(users)
	nextId := 0
	for conn := range c {
		go authConnection(nextId, conn, users)
		nextId++
	}
}

func authConnection(id int, conn Connection, users chan<- ActiveUser) {
	for {
		conn.Write <- "By what alias would you like to be known?\n"
		conn.Prompt <- true
		nick, ok := <-conn.Read
		if !ok {
			// Client disconnected.
			return
		}
		nick = strings.TrimSpace(nick)
		if nick != "" {
			conn.Write <- fmt.Sprintf("Welcome to a Science Fiction Universe, %v!\n", nick)
			users <- ActiveUser{id, nick, conn, time.Now()}
			return
		}
		conn.Write <- "Bad answer...\n"
	}
}
