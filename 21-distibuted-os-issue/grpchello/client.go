package grpchello

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func startGrpcHelloClient() {
	conn, _ := grpc.Dial("127.0.0.1:3000", grpc.WithInsecure())
	defer conn.Close()
	stub := NewHelloServiceClient(conn)
	result, _ := stub.Say(context.Background(), &Person{Name: "zhangsan"}, grpc.EmptyCallOption{})
	fmt.Println(result.Text)
}
