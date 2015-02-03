package core

import (
	"fmt"
	"strings"
	"time"
	"unicode"
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
		conn.Write <- "What is your nom de plume?\n"
		nick, ok := <-conn.Read
		if !ok {
			// Client disconnected.
			return
		}
		nick = strings.TrimFunc(nick, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsPunct(r)
		})
		if nick != "" && len(nick) < 16 {
			conn.Write <- fmt.Sprintf("Welcome to a Science Fiction Universe, %v!\n", nick)
			users <- ActiveUser{id, nick, conn, time.Now()}
			return
		}
		conn.Write <- "Bad answer...\n"
	}
}
