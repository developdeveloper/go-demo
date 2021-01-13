package hellorpc

import (
	"fmt"
	"net/rpc"
)

func startHelloRpcClient() {
	var reply string
	client, _ := rpc.Dial("tcp", "127.0.0.1:3000")
	client.Call("HelloService.Say", "zhangsan", &reply)

	//异步
	//done := client.Go("HelloService.Say", "zhangsan", &reply, nil).Done
	//<-done

	fmt.Println(reply)
}
