package main

import (
    "flag"
    "fmt"
    "github.com/gnarlyskier/wander"
)

// Debug handler for clients that simply echos.
func debugClientHandler(c <-chan wander.Client) {
    for client := range c {
        // Echo lines and handle exit commands.
        go func() {
            for s := range client.Read {
                client.Write <- s + "\n"
                fmt.Printf("read: '%v'\n", s)
                if (s == "exit") {
                    fmt.Println("dropped client on request")
                    client.Close()
                }
            }
        }()
    }
}

func main() {
    port := flag.Uint("port", 4000, "port on which to listen for connections")
    flag.Parse()

    // Handle connected clients.
    c := make(chan wander.Client)
    go debugClientHandler(c)

    if err := wander.ServeForever(*port, c); err != nil {
        fmt.Printf("Failed start server: %v", err)
    }
}
