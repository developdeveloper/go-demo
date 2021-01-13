package betterrpc

import (
	"fmt"
)

func startBetterRpcClient() {
	var reply string
	client, _ := DialHelloService("tcp", "localhost:3000")
	client.Hello("zhangsan", &reply)

	//异步
	//done := client.AsyncHello("zhangsan", &reply)
	//<-done

	fmt.Println(reply)
}
