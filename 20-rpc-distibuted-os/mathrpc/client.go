package mathrpc

import (
	"fmt"
	"net/rpc"
)

func startMathRpcClient() {
	var reply int
	client, _ := rpc.Dial("tcp", "127.0.0.1:3000")

	expr := Expr{"add", 1, 2}
	client.Call("MathService.Calc", expr, &reply)
	fmt.Println(reply)

	expr = Expr{"mul", 2, 3}
	client.Call("MathService.Calc", expr, &reply)
	fmt.Println(reply)
}
