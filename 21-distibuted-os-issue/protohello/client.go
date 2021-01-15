package protohello

import (
	"fmt"
	"net/rpc"
)

func startHelloProtoRpcClient() {
	var reply String
	client, _ := rpc.Dial("tcp", "127.0.0.1:3000")
	client.Call("HelloService.Say", String{Value: "zhangsan"}, &reply)
	fmt.Println(reply.Value)
}
