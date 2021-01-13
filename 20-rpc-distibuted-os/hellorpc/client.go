package hellorpc

import (
	"fmt"
	"net/rpc"
)

func startHelloRpcClient() {
	var reply string
	client, _ := rpc.Dial("tcp", "127.0.0.1:3000")
	client.Call("HelloService.Say", "zhangsan", &reply)
	fmt.Println(reply)
}
