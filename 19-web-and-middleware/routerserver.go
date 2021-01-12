package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func startRouterServer() {
	router := httprouter.New()
	router.GET("/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintln(w, fmt.Sprintf("%s 你好!", ps.ByName("name")))
	})
	http.ListenAndServe(":3000", router)
}
