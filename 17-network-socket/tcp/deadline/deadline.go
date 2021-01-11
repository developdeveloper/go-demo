package deadline

import (
	"fmt"
	"net"
	"time"
)

type client struct {
	conn    net.Conn
	timeout int
}

func handle(c *client) {
	buf := make([]byte, 100)
	for {
		size, err := c.conn.Read(buf)
		if err != nil || size == 0 {
			fmt.Println(err)
			c.conn.Close()
			return
		}

		c.conn.SetDeadline(time.Now().Add(3 * time.Second))
		fmt.Println(string(buf))
	}
}

func startServerWithTimeout() {
	listener, _ := net.Listen("tcp", "127.0.0.1:3000")
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		conn.SetDeadline(time.Now().Add(3 * time.Second))
		c := &client{conn, 3}
		go handle(c)
	}
}
