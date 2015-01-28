package core

import (
    "strconv"
    "time"
)

type ActiveUser struct {
    Id int
    Conn Connection
    timestamp time.Time
}

func AuthNewUsers(c <-chan Connection, users chan<- ActiveUser) {
    defer func() {
        close(users)
    }()
    for conn := range c {
        go authConnection(conn, users)
    }
}

func authConnection(conn Connection, users chan<- ActiveUser) {
    conn.Write <- "Welcome to the Napoleonic Wars!\n"
    for {
        conn.Write <- "What is your client id?\n"
        id, err := strconv.Atoi(conn.Prompt())
        if err == nil {
            users <- ActiveUser{id, conn, time.Now()}
            return
        }
        conn.Write <- "Bad answer...\n"
    }
}
