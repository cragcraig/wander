package main

import (
    "flag"
    "fmt"
    "github.com/gnarlyskier/wander/core"
)

// Debug handler for clients that simply echos.
func debugUserHandler(c <-chan core.ActiveUser) {
    for user := range c {
        // Echo lines and handle exit commands.
        go func() {
            for s := range user.Client.Read {
                user.Client.Write <- s + "\n"
                fmt.Printf("read: '%v'\n", s)
                if (s == "exit") {
                    fmt.Println("dropped user on request")
                    user.Client.Close()
                }
            }
        }()
    }
}

func main() {
    port := flag.Uint("port", 4000, "port on which to listen for connections")
    flag.Parse()

    clients := make(chan core.Client)
    users := make(chan core.ActiveUser)
    go core.AuthNewUsers(clients, users)

    // Handle connected users.
    go debugUserHandler(users)

    if err := core.ServeForever(*port, clients); err != nil {
        fmt.Printf("Failed start server: %v", err)
    }
}
