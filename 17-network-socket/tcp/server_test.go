package tcp

import "testing"

func Test_startTcpServer(t *testing.T) {
	startTcpServer()
}

func Test_startTcpClient(t *testing.T) {
	startTcpClient()
}

func Test_acceptBacklog(t *testing.T) {
	go startTcpServer2()
	startTcpClient2()
}

func Test_dialTimeout(t *testing.T) {
	go startTcpServer3()
	startTcpClient3()
}
