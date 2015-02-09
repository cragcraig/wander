package main

import (
	"flag"
	"fmt"
	"github.com/gnarlyskier/wander/ascii"
	"github.com/gnarlyskier/wander/core"
)

func handleConnections(conn <-chan *core.Connection) {
	for c := range conn {
		buf := make([]byte, 80*24)
		for i := range buf {
			buf[i] = 'o'
		}
		c.RawWrite <- ascii.Render(buf, 80, 24)
	}
}

func main() {
	port := flag.Uint("port", 4000, "port on which to listen for connections")
	flag.Parse()

	conns := make(chan *core.Connection)
	go handleConnections(conns)

	if err := core.ServeForever(*port, conns); err != nil {
		fmt.Printf("Failed start server: %v", err)
	}
}
