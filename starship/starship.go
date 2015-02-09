package main

import (
	"flag"
	"fmt"
	"github.com/gnarlyskier/wander/ascii"
	"github.com/gnarlyskier/wander/core"
	"github.com/gnarlyskier/wander/verse"
)

func handleConnections(conn <-chan *core.Connection) {
	v := verse.SimpleVerse{[]verse.PosRenderable{
		verse.PosRenderable{5, 5, ascii.TieFighter}}}
	buf := ascii.CreateBuffer(80, 24)
	for c := range conn {
		v.Render(0, 0, buf)
		c.RawWrite <- buf.Render()
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
