package betterrpc

import (
	"testing"
	"time"
)

func Test_startBetterRpcClient(t *testing.T) {
	go startBetterRpcServer()
	time.Sleep(1 * time.Second)
	startBetterRpcClient()
}
