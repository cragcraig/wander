package main

import (
	"flag"
	"fmt"
	"github.com/cragcraig/wander/core"
	"github.com/cragcraig/wander/env"
	"github.com/cragcraig/wander/nlp"
	"github.com/cragcraig/wander/verbs"
)

func userRoom(user *core.ActiveUser, actions chan<- *env.Action) {
	player := env.CreatePlayer(user)
	actions <- player.CreateAction(verbs.EnterRoom, nil, nil, nil)
	defer func() {
		actions <- player.CreateAction(verbs.LeaveRoom, nil, nil, nil)
	}()
	for input := range user.Conn.Read {
		action := nlp.ParsePlayerAction(player, input)
		if action != nil {
			actions <- action
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
