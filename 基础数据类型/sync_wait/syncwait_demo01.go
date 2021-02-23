package sync_wait

import (
	"fmt"
	"github.com/golang/groupcache/singleflight"
	"sync"
	"time"
)

/**
*	sync.WaitGroup例子
 */
var syncWait sync.WaitGroup
var mu sync.Mutex
var m map[string]int

func NewDelayReturn(dur time.Duration, n int) func() (interface{}, error) {
	return func() (interface{}, error) {
		fmt.Println("计算逻辑执行次数...")
		time.Sleep(dur) //计算逻辑
		return n, nil
	}
}
func SyncWait() {
	g := singleflight.Group{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		fmt.Println("执行我了1")
		ret, err := g.Do("key", NewDelayReturn(time.Second*1, 1))
		if err != nil {
			panic(err)
		}
		fmt.Printf("key-1 get %v\n", ret)
		wg.Done()
	}()
	//go func() {
	//	time.Sleep(100 * time.Millisecond) // make sure this is call is later
	//	fmt.Println("执行我了2")
	//	ret, err := g.Do("key", NewDelayReturn(time.Second*3, 2))
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("key-2 get %v\n", ret)
	//	wg.Done()
	//}()
	go func() {
		fmt.Println("执行我了key1 1")
		ret, err := g.Do("key1", NewDelayReturn(time.Second*3, 2))
		if err != nil {
			panic(err)
		}
		fmt.Printf("key-1-1 get %v\n", ret)
		wg.Done()
	}()
	//go func() {
	//	fmt.Println("执行我了key1 2")
	//	ret, err := g.Do("key1", NewDelayReturn(time.Second*3, 2))
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("key-1-2 get %v\n", ret)
	//	wg.Done()
	//}()
	wg.Wait()
}

//func SyncWait() {
//	go calculate("a", 0)
//	go calculate("a", 1)
//	go calculate("a", 2)
//	time.Sleep(4 * time.Second)
//	fmt.Printf("%+v", m)
//
//}

//func calculate(k string, v int) (rs int) {
//	mu.Lock()
//	if v, ok := m[k]; ok {
//		mu.Unlock()
//		syncWait.Wait() //等待其它结果返回
//		return v
//	}
//	if m == nil {
//		m = make(map[string]int)
//	}
//	syncWait.Add(1)
//	time.Sleep(1 * time.Second) //假如是计算结果的任务
//	m[k] = v
//	mu.Unlock()
//	syncWait.Done()
//	return v
//}
