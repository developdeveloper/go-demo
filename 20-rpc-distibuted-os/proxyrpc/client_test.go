package proxyrpc

import "testing"

func Test_startProxyRpcClient(t *testing.T) {
	go startProxyRpcServer()
	startProxyRpcClient()
}
