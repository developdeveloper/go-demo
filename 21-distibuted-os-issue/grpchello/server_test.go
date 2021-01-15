package grpchello

import "testing"

func Test_startGrpcHelloServer(t *testing.T) {
	startGrpcHelloServer()
	select {}
}
