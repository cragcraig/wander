package main

import (
    "bytes"
    "fmt"
    "io"
    "net"
    "strconv"
)

var PORT int = 4000
var BUFSIZE int = 512

type Client struct {
    Conn net.Conn
    Read chan string
    Write chan string
}

func (c *Client) Close() {
    c.Conn.Close()
}

func readLines(m chan<- string, r io.Reader) {
    defer func() {
        close(m)
    }()
    var buffer bytes.Buffer
    b := make([]byte, BUFSIZE)
    for {
        n, err := r.Read(b)
        if err != nil {
            return
        }
        for i := 0; i < n; i++ {
            c := b[i]
            if strconv.IsPrint(rune(c)) {
                buffer.WriteByte(c)
            } else if (buffer.Len() != 0) {
                m <- buffer.String()
                buffer.Reset()
            }
        }
    }
}

func writeLines(m <-chan string, w io.Writer) {
    for s := range m {
        w.Write([]byte(s))
    }
}

func detachConnection(conn net.Conn) Client {
    r := make(chan string)
    go readLines(r, conn)
    w := make(chan string)
    go writeLines(w, conn)
    return Client{conn, r, w}
}

func acceptConnectionsForever(ln net.Listener, c chan<- Client) {
    for {
        if conn, err := ln.Accept(); err == nil {
            fmt.Println("Client connected")
            c <- detachConnection(conn)
        }
    }
}

// Debug handler for clients that simply echos.
func debugClientHandler(c <-chan Client) {
    for client := range c {
        // Echo lines and handle exit commands.
        go func() {
            for s := range client.Read {
                client.Write <- s + "\n"
                fmt.Printf("read: '%v'\n", s)
                if (s == "exit") {
                    fmt.Println("closed client on request")
                    client.Close()
                }
            }
        }()
    }
}

func main() {
    if ln, err := net.Listen("tcp", fmt.Sprintf(":%v", PORT)); err != nil {
        fmt.Println("Failed to listen", err)
    } else {
        fmt.Printf("Listening on localhost:%v\n", PORT)
        c := make(chan Client)
        // Handle connected clients.
        go debugClientHandler(c)
        acceptConnectionsForever(ln, c)
    }
}
