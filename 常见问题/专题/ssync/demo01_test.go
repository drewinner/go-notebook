package ssync

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestDemo01(t *testing.T) {
	//Demo01()
	//Demo02Call()
	fmt.Println(1 << 0)
	var i int32 = 20
	fmt.Println(i)
	b := atomic.CompareAndSwapInt32(&i, 20, 2)
	fmt.Println(b)
	fmt.Println(i)
}
