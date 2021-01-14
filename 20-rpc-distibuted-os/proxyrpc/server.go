package proxyrpc

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

func startProxyRpcServer() {
	rpc.Register(new(HelloService))

	for {
		// 反过来拨号到外网的 ip 地址上
		conn, err := net.Dial("tcp", "127.0.0.1:3000")
		if err == nil {
			// 因为外网客户端还未监听报错
			continue
		}

		rpc.ServeConn(conn)
		conn.Close()
	}
}
