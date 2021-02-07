package chap03

import (
	"fmt"
	"net/http"
)

func handler07(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}

func Demo07Test() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("root direct...")
	})
	http.HandleFunc("/hello/", handler07)

	_ = http.ListenAndServe(":80", nil)
}
