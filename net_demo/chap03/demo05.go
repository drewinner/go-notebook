package chap03

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello05(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func log05(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}
func Demo05Test() {
	server := http.Server{
		Addr: ":80",
	}
	http.HandleFunc("/hello", log05(hello05))
	server.ListenAndServe()
}
