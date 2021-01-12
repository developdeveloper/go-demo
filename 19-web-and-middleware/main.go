package main

import (
	"net/http"
)

func main() {
	// 作用在 DefaultServeMux
	// http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {})

	// 作用在 DefaultServeMux
	http.ListenAndServe(":3000", nil)
}
