package web

import (
	"fmt"
	"net/http"
	"time"
)

func startHttpServer() {
	http.HandleFunc("/time", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(time.Now().String()))
	})

	http.HandleFunc("/headers", func(writer http.ResponseWriter, request *http.Request) {
		for key, headers := range request.Header {
			for _, val := range headers {
				fmt.Fprintf(writer, "%v:%v\n", key, val)
			}
		}
	})

	http.ListenAndServe(":3000", nil)
}
