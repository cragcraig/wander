package core

import (
    "strconv"
    "time"
)

type ActiveUser struct {
    ClientId int
    Client Client
    timestamp time.Time
}

func AuthNewUsers(c <-chan Client, conns chan<- ActiveUser) {
    defer func() {
        close(conns)
    }()
    for cl := range c {
        go authClient(cl, conns)
    }
}

func authClient(cl Client, conns chan<- ActiveUser) {
    cl.Write <- "Welcome to the Napoleonic Wars!\n"
    for {
        cl.Write <- "What is your client id?\n"
        id, err := strconv.Atoi(cl.Prompt())
        if err == nil {
            conns <- ActiveUser{id, cl, time.Now()}
            return
        }
        cl.Write <- "Bad answer...\n"
    }
}
