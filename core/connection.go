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
}

func (conn *Connection) Prompt() (string, bool) {
	conn.Write <- "> "
	v, ok := <-conn.Read
	return v, ok
}

func (conn *Connection) Close() {
	conn.netConn.Close()
}

func readLines(c chan<- string, r io.Reader) {
	defer func() {
		close(c)
	}()
	var buffer bytes.Buffer
	b := make([]byte, BUFSIZE)
	for {
		n, err := r.Read(b)
		if err != nil {
			return
		}
		for i := 0; i < n; i++ {
			next := b[i]
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
		w.Write([]byte(s))
	}
}

func detachConnection(conn net.Conn) Connection {
	r := make(chan string)
	go readLines(r, conn)
	w := make(chan string)
	go writeLines(w, conn)
	return Connection{conn, r, w}
}

func acceptConnectionsForever(ln net.Listener, c chan<- Connection) {
	for {
		if conn, err := ln.Accept(); err == nil {
			fmt.Println("Connection connected")
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
