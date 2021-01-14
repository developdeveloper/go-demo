package hellorpc

import (
	"fmt"
	"net/rpc"
)

func startHelloRpcClient() {
	var reply string
	client, _ := rpc.Dial("tcp", "127.0.0.1:3000")

	//使用json编码
	//conn, _ := net.Dial("tcp", "127.0.0.1:3000")
	//client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	client.Call("HelloService.Say", "zhangsan", &reply)

	//异步
	//done := client.Go("HelloService.Say", "zhangsan", &reply, nil).Done
	// 继续做其它的事情
	//<-done

	fmt.Println(reply)
}
