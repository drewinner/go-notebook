package chap01

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "hello world :%s", request.URL.Path[1:])
}

func handler01(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "hello world 01 :%s", request.URL.Path[1:])
}

func handler02(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "hello world 02:%s", request.URL.Path[1:])
}

func handler03(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "hello world 03:%s", request.URL.Path[1:])
}

func Demo01Test() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/aa/", handler01)
	http.HandleFunc("/aa/bb/", handler02)
	http.HandleFunc("/aa/bb/cc", handler03)
	_ = http.ListenAndServe(":8080", nil)
}
