package hellorpc

import (
	"testing"
	"time"
)

func Test_startHelloRpcClient(t *testing.T) {
	go startHelloRpcServer()
	time.Sleep(1 * time.Second)
	startHelloRpcClient()
}
