package mathrpc

import (
	"net"
	"net/rpc"
)

//MathService 打招呼的服务
type MathService struct{}

func (ms *MathService) Calc(expr Expr, reply *int) error {
	switch expr.Method {
	case "add":
		*reply = expr.Left + expr.Right
	case "mul":
		*reply = expr.Left * expr.Right
	}
	return nil
}

func startMathRpcServer() {
	rpc.RegisterName("MathService", new(MathService))
	listener, _ := net.Listen("tcp", ":3000")
	conn, _ := listener.Accept()
	rpc.ServeConn(conn)
}
