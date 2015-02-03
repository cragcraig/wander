package core

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"
)

var BUFSIZE int = 512

type Connection struct {
	netConn net.Conn
	Read    <-chan string
	Write   chan<- string
	Prompt  chan<- bool
}

func (conn *Connection) Close() {
	fmt.Printf("Closed connection to %v\n", conn.netConn.RemoteAddr())
	conn.netConn.Close()
}

func readLines(c chan<- string, r io.Reader) {
	defer close(c)
	var buffer bytes.Buffer
	b := make([]byte, BUFSIZE)
	for {
		n, err := r.Read(b)
		if err != nil {
			return
		}
		for i := 0; i < n; i++ {
			next := b[i]
            // TODO(craig): Handle backspace
			if strconv.IsPrint(rune(next)) {
				buffer.WriteByte(next)
			} else if buffer.Len() != 0 {
				c <- buffer.String()
				buffer.Reset()
			}
		}
	}
}

func writeLines(c <-chan string, w io.Writer) {
	for s := range c {
        // TODO(craig): Backspace over prompt and characters in buffer
		w.Write([]byte(s))
	}
}

func writePrompts(p <-chan bool, w chan<- string) {
	for _ = range p {
        // TODO(craig): Print contents of buffer (if typing was interrupted)
		w <- "> "
	}
}

func detachConnection(conn net.Conn) Connection {
	r := make(chan string)
	go readLines(r, conn)
	w := make(chan string)
	go writeLines(w, conn)
	p := make(chan bool)
	go writePrompts(p, w)
	return Connection{conn, r, w, p}
}

func acceptConnectionsForever(ln net.Listener, c chan<- Connection) {
	for {
		if conn, err := ln.Accept(); err == nil {
			fmt.Printf("Accepted new connection from %v\n", conn.RemoteAddr())
			c <- detachConnection(conn)
		}
	}
}

func ServeForever(port uint, c chan<- Connection) error {
	if ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port)); err != nil {
		return err
	} else {
		fmt.Printf("Listening on port %v\n", port)
		acceptConnectionsForever(ln, c)
		return nil
	}
}
