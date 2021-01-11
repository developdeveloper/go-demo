package udp

import "testing"

func Test_startUdpServer(t *testing.T) {
	go startUdpServer()
	startUdpClient()
}
