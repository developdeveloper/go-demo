package hellorpc

import (
	"net"
	"net/rpc"
)

//HelloService 打招呼的服务
type HelloService struct{}

func (hs *HelloService) Say(name string, reply *string) error {
	*reply = "hi, " + name
	return nil
}

func startHelloRpcServer() {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, _ := net.Listen("tcp", ":3000")
	conn, _ := listener.Accept()
	rpc.ServeConn(conn)
}
