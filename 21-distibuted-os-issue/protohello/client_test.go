package protohello

import (
	"testing"
	"time"
)

func Test_startHelloProtoRpcClient(t *testing.T) {
	go startHelloProtoRpcServer()
	time.Sleep(1 * time.Second)
	startHelloProtoRpcClient()
}
