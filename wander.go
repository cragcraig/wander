package main

import (
	"flag"
	"fmt"
	"github.com/gnarlyskier/wander/core"
)

// Debug handler for users that simply echos.
func debugUserHandler(c <-chan core.ActiveUser) {
	for user := range c {
		// Echo lines and handle exit commands.
		go func() {
			for s := range user.Conn.Read {
				user.Conn.Write <- s + "\n"
				fmt.Printf("%v: %v\n", user.Id, s)
				if s == "exit" {
					fmt.Println("dropped user on request")
					user.Conn.Close()
				}
			}
		}()
	}
}

func main() {
	port := flag.Uint("port", 4000, "port on which to listen for connections")
	flag.Parse()

	conns := make(chan core.Connection)
	users := make(chan core.ActiveUser)
	go core.AuthNewUsers(conns, users)

	// Handle connected users.
	go debugUserHandler(users)

	if err := core.ServeForever(*port, conns); err != nil {
		fmt.Printf("Failed start server: %v", err)
	}
}
