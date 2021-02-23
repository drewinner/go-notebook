package singleflight

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func NewDelayReturn(dur time.Duration, n int) func() (interface{}, error) {
	return func() (interface{}, error) {
		fmt.Println("计算逻辑执行次数...")
		time.Sleep(dur) //计算逻辑
		return n, nil
	}
}

type person struct {
	name string
}

func TestGroup_Do(t *testing.T) {

	g := Group{}
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		fmt.Println("执行我了1")
		ret, err := g.Do("key", NewDelayReturn(time.Second*10, 1))
		if err != nil {
			panic(err)
		}
		fmt.Printf("key-1 get %v\n", ret)
		wg.Done()
	}()
	go func() {
		time.Sleep(100 * time.Millisecond) // make sure this is call is later
		fmt.Println("执行我了2")
		ret, err := g.Do("key", NewDelayReturn(time.Second*3, 2))
		if err != nil {
			panic(err)
		}
		fmt.Printf("key-2 get %v\n", ret)
		wg.Done()
	}()
	go func() {
		fmt.Println("执行我了3")
		ret, err := g.Do("key", NewDelayReturn(time.Second*3, 3))
		if err != nil {
			panic(err)
		}
		fmt.Printf("key-3 get %v\n", ret)
		wg.Done()
	}()
	wg.Wait()
}
