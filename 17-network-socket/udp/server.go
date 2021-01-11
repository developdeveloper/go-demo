package udp

import (
	"fmt"
	"net"
)

func startUdpServer() {
	listener, _ := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 3000})
	buf := make([]byte, 100)
	for {
		size, remoteAddr, err := listener.ReadFromUDP(buf)
		if err == nil && size > 0 {
			fmt.Println(remoteAddr.String())
			fmt.Println(string(buf))
		}
	}
}

func startUdpClient() {
	ip := net.ParseIP("127.0.0.1")
	// 自动分配
	// local := &net.UDPAddr{IP: net.IPv4zero, Port: 1110}
	local := &net.UDPAddr{IP: ip, Port: 3001}
	remote := &net.UDPAddr{IP: ip, Port: 3000}
	conn, _ := net.DialUDP("udp", local, remote)
	defer conn.Close()
	conn.Write([]byte("can you hear me?\n"))
}
