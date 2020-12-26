package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	test()
}
func test() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	handler(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handler(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
func test02() {
	var c = make(chan int)
	go func(chan int) {
		time.Sleep(time.Second * 1)
		c <- 1
	}(c)
	select {
	case <-c:
		fmt.Println("receive c ...")
	}
	time.Sleep(1 * time.Second)
}
