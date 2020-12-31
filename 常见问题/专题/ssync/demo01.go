package ssync

import (
	"fmt"
	"sync"
	"time"
)

func Demo01() {
	var l sync.Mutex
	a := 1
	go func(mutex sync.Mutex) {
		l.Lock()
		fmt.Println("aaaaaa")
		time.Sleep(10 * time.Second)
		l.Unlock()
	}(l)
	time.Sleep(1 * time.Second)
	l.Lock()
	a++
	fmt.Println("bbbb")
	l.Unlock()
	time.Sleep(100 * time.Second)
}
