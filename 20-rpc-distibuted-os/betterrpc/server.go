package betterrpc

import (
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(name string, reply *string) error {
	*reply = "hi, " + name
	return nil
}

func startBetterRpcServer() {
	RegisterHelloService(new(HelloService))
	listener, _ := net.Listen("tcp", ":3000")

	//使用 for 循环服务多个客户端
	conn, _ := listener.Accept()
	rpc.ServeConn(conn)
}
