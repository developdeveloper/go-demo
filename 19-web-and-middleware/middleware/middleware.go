package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// 参数和返回用 HandlerFunc 也可以，但是 http.Handler 是接口更通用
// type middleware func(handler http.Handler) http.Handler
func loggerMiddleware(real http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "before "+time.Now().String())
		real.ServeHTTP(writer, request)
		fmt.Fprintln(writer, "\nafter "+time.Now().String())
	})
}

func startMiddlewareServer() {
	serve1 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("我前后有时间"))
	}
	// 包装一下，让它具有 ServeHTTP 函数申明
	handler := http.HandlerFunc(serve1)
	http.Handle("/1", loggerMiddleware(handler))

	serve2 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("我前后也有时间"))
	}
	http.Handle("/2", loggerMiddleware(http.HandlerFunc(serve2)))
	//http.HandleFunc("/2", loggerMiddleware(http.HandlerFunc(serve2)).ServeHTTP)

	http.ListenAndServe(":3000", nil)
}
