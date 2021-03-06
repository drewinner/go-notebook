package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

var datas03 []string

func init() {
	runtime.SetMutexProfileFraction(1)
}
func main() {
	var m sync.Mutex
	var datas03 = make(map[int]struct{})
	for i := 0; i < 999; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			time.Sleep(10 * time.Second)
			datas03[i] = struct{}{}
		}(i)
	}
	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}
