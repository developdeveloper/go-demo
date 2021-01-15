package protohello

import (
	"net"
	"net/rpc"
)

func startHelloProtoRpcServer() {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, _ := net.Listen("tcp", ":3000")
	conn, _ := listener.Accept()
	rpc.ServeConn(conn)
}
