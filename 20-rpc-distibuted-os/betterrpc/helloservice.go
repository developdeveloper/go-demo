package betterrpc

import "net/rpc"

const HelloServiceName = "20-rpc-distibuted-os/betterrpc/HelloService"

// Server Func
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

// Client Func
type HelloServiceClient struct {
	*rpc.Client
}

func (p *HelloServiceClient) Hello(name string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", name, reply)
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, _ := rpc.Dial(network, address)
	return &HelloServiceClient{Client: c}, nil
}
