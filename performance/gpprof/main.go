package main

import (
	"log"
	"net/http"
	"time"
	_ "net/http/pprof"
)

var datas []string

func main() {
	go func(){
		for {
			log.Printf("len: %d",add("go-programming-tour-book"))
			time.Sleep(time.Second*10)
		}
	}()
	_ = http.ListenAndServe("0.0.0.0:6060",nil)
}

func add(str  string) int{
	data := []byte(str)
	datas = append(datas,string(data))
	return len(datas)
}