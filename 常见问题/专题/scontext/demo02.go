package scontext

import (
	"context"
	"fmt"
	"time"
)

func Demo02() {
	messages := make(chan int, 10)
	for i := 0; i < 10; i++ {
		messages <- i
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			select {
			case <-ctx.Done():
				go func(ctx context.Context) {
					for {
						fmt.Println("child process 1111 ......")
						time.Sleep(1 * time.Second)
						select {
						case <-ctx.Done():
							fmt.Println("child process 1 interrupt....")
							return
						}
					}
				}(ctx)
				fmt.Println("child process interrupt....")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}(ctx)
	defer close(messages)
	defer cancel()
	select {
	case <-ctx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("main process exit!")
	}
}
