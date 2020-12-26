package scontext

import (
	"context"
	"errors"
	"fmt"
)

var c = 0

func WithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	speakMemo(ctx, cancel)
}

func speakMemo(context context.Context, cancelFunc context.CancelFunc) {
	for {
		select {
		case <-context.Done():
			fmt.Println("done...")
			return
		default:
			fmt.Println("exec default func ...")
			err := doSome(3)
			if err != nil {
				fmt.Println("cancelFunc()")
				cancelFunc()
			}
		}
	}
}

func doSome(i int) error {
	c++
	if c > i {
		return errors.New("err occur")
	}
	return nil
}
