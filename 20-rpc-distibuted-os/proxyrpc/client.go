package proxyrpc

import (
	"fmt"
	"net"
	"net/rpc"
)

func startProxyRpcClient() {
	var reply string

	// 外网的客户端主动提供 tcp 服务等待连接
	listener, _ := net.Listen("tcp", ":3000")
	conn, _ := listener.Accept()

	// 构建 rpc 客户端对象
	client := rpc.NewClient(conn)
	defer client.Close()
	client.Call("HelloService.Say", "zhangsan", &reply)
	fmt.Println(reply)
}
