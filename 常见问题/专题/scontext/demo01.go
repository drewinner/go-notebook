package scontext

import (
	"fmt"
	"time"
)

func Demo01() {
	messages := make(chan int, 5)
	done := make(chan bool)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			select {
			case b := <-done:
				fmt.Printf("child process interrupt...%+v\n", b)
				return
			case m := <-messages:
				fmt.Printf("send message: %d\n", m)
			}
		}
	}()
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			select {
			case b := <-done:
				fmt.Printf("child process1 interrupt...%+v\n", b)
				return
			case m := <-messages:
				fmt.Printf("send message1: %d\n", m)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		messages <- i
	}
	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(100 * time.Second)
	fmt.Println("Demo01 process exit!")
}
