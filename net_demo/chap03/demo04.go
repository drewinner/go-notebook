package chap03

import (
	"fmt"
	"net/http"
)

/**
*	处理器函数
 */
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func Demo04Test() {
	server := http.Server{
		Addr: ":80",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	_ = server.ListenAndServe()
}
