package grpchello

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

type HelloServiceServerImpl struct{}

func (hssi *HelloServiceServerImpl) Say(ctx context.Context, args *Person) (*Result, error) {
	reply := &Result{Text: "hi, " + args.Name}
	return reply, nil
}

func startGrpcHelloServer() {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(HelloServiceServerImpl))
	conn, _ := net.Listen("tcp", ":3000")
	grpcServer.Serve(conn)
}
