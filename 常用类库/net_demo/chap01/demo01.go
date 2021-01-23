package chap01

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "hello world :%s", request.URL.Path[1:])
}

type a string

// ServeHTTP calls f(w, r).
func (f a) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w.Header())
}

func Demo01Test() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/aa/", handler)
	http.HandleFunc("/aa/bb/", handler)
	http.HandleFunc("/aa/bb/cc", handler)
	http.Handle("aaa", a("www"))
	_ = http.ListenAndServe(":8080", nil)
}
