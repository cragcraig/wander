package main

import (
	"flag"
	"fmt"
	"github.com/gnarlyskier/wander/core"
	"github.com/gnarlyskier/wander/env"
	"github.com/gnarlyskier/wander/simpleverbs"
)

func userRoom(user *core.ActiveUser, actions chan<- *env.Action) {
	actions <- env.EnterRoom.CreateUserAction(user, nil, nil, nil)
	defer func() {
		actions <- env.LeaveRoom.CreateUserAction(user, nil, nil, nil)
	}()
	for cmd := range user.Conn.Read {
		switch cmd {
		case "exit":
			user.Conn.Close()
		default:
			actions <- simpleverbs.Talk.CreateUserAction(user, nil, nil, []string{cmd})
		}
	}
}

func demoRoom(c <-chan *core.ActiveUser, room *env.Room) {
	for user := range c {
		go userRoom(user, room.Actions)
	}
}

func main() {
	port := flag.Uint("port", 4000, "port on which to listen for connections")
	flag.Parse()

	conns := make(chan *core.Connection)
	users := make(chan *core.ActiveUser)
	go core.AuthNewUsers(conns, users)

	// Handle connected users.
	go demoRoom(users, env.CreateRoom())

	if err := core.ServeForever(*port, conns); err != nil {
		fmt.Printf("Failed start server: %v", err)
	}
}
