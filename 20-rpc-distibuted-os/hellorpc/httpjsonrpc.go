package hellorpc

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func startHttpJsonRpcServer() {
	rpc.RegisterName("HelloService", new(HelloService))

	http.HandleFunc("/say", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			writer,
			request.Body,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	// curl localhost:3000/say -X POST --data '{"method":"HelloService.Say","params":["zhangsan"],"id":0}'
	http.ListenAndServe(":3000", nil)
}
