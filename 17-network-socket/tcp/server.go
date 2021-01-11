package tcp

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	str := fmt.Sprintf("%s\n", time.Now().Format("2006-01-02 15:04:05"))
	conn.Write([]byte(str))
}

func startTcpServer() {
	listener, _ := net.Listen("tcp", "127.0.0.1:3000")
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go handleClient(conn)
	}
}

func startTcpServer2() {
	listener, _ := net.Listen("tcp", "127.0.0.1:3000")
	defer listener.Close()
	for {
		time.Sleep(5 * time.Second)
		conn, _ := listener.Accept()
		log.Println("Accept at " + time.Now().Format("15:04:05"))
		go handleClient(conn)
	}
}

func startTcpServer3() {
	listener, _ := net.Listen("tcp", "127.0.0.1:3000")
	defer listener.Close()
	select {}
}

func startTcpClient() {
	conn, _ := net.Dial("tcp", "127.0.0.1:3000")
	buf := make([]byte, 100)
	conn.Read(buf)
	fmt.Println(string(buf))
}

func startTcpClient2() {
	for i := 0; i < 1024; i++ {
		conn, _ := net.Dial("tcp", "127.0.0.1:3000")
		log.Println(conn.LocalAddr().String() + " dial at " + time.Now().Format("15:04:05"))
	}
}

func startTcpClient3() {
	for i := 0; i < 1024; i++ {
		_, err := net.Dial("tcp", "127.0.0.1:3000")
		// 设置超时时间
		// _, err := net.DialTimeout("tcp", "127.0.0.1:3000", 1*time.Second)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("dial at " + time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}
