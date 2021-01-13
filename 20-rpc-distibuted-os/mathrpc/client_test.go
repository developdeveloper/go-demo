package mathrpc

import (
	"testing"
	"time"
)

func Test_startMathRpcClient(t *testing.T) {
	go startMathRpcServer()
	time.Sleep(1 * time.Second)
	startMathRpcClient()
}
