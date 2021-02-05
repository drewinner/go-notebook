package channel_demo

import (
	"fmt"
	"math"
	"runtime"
)

func ChannelDemoDemo01() {
	runtime.GOMAXPROCS(1)
	//debug.SetMaxThreads(1)
	//ch := make(chan int, 3)
	//runtime.GOMAXPROCS(1)
	//debug.SetMaxThreads(1)
	for i := 0; i < math.MaxInt32; i++ {
		go func() {
			fmt.Println("world", i)
		}()
	}
}
