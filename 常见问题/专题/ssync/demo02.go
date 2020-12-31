package ssync

import (
	"fmt"
	"sync"
	"time"
)

func Demo02Call() {
	//设置访问路由
	var p = new(sync.Mutex)
	go func1(p, "A1")
	go func1(p, "A2")
	go func1(p, "A3")

	time.Sleep(time.Second * 36)
}
func func1(s *sync.Mutex, curNum string) {
	fmt.Println("I'm ", curNum)
	s.Lock() //加锁
	fmt.Println(curNum, "------begin--------")
	fmt.Println(curNum, "已经加锁")
	fmt.Println("这部分的代码，同一时刻，只让一个G来执行")
	time.Sleep(15 * time.Second)
	fmt.Println(curNum, "准备解锁")
	fmt.Println(curNum, "------end--------")
	s.Unlock() //执行结束时候解锁
}
