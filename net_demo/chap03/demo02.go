package chap03

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func Demo02Test() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:80",
		Handler: &handler,
	}
	_ = server.ListenAndServe()
}
