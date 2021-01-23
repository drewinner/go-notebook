package chap03

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct {
}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}
func Demo03Test() {
	hello := HelloHandler{}
	world := WorldHandler{}
	server := http.Server{
		Addr: ":80",
	}
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	_ = server.ListenAndServe()
}
