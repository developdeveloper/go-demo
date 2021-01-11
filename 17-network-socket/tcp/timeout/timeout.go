package timeout

import (
	"fmt"
	"net"
	"time"
)

type client struct {
	conn       net.Conn
	activeChan chan struct{}
	timeout    int
}

func handle(c *client) {
	for {
		buf := make([]byte, 100)
		size, err := c.conn.Read(buf)
		if err == nil || size > 0 {
			c.activeChan <- struct{}{}
			fmt.Println(string(buf))
		}
	}
}

func alive(c *client) {
	for {
		select {
		case <-c.activeChan:
			fmt.Println("keep alive")
		case <-time.After(time.Duration(c.timeout) * time.Second):
			c.conn.Write([]byte("goodbye\n"))
			c.conn.Close()
		}
	}
}

func startServerWithTimeout() {
	listener, _ := net.Listen("tcp", "127.0.0.1:3000")
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		c := &client{conn, make(chan struct{}), 3}
		go handle(c)
		go alive(c)
	}
}
