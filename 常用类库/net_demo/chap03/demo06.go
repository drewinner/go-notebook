package chap03

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

type HelloHandler06 struct {
}

func (h *HelloHandler06) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func log06(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("protect ....")
		h.ServeHTTP(w, r)
	})
}

func Demo06Test() {
	server := http.Server{
		Addr: ":80",
	}
	hello := &HelloHandler06{}
	http.Handle("/hello", protect(log06(hello)))
	server.ListenAndServe()
}
