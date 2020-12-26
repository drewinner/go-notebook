package main

import (
	"fmt"
	"sync"
)

func main() {
	var m = map[int]int{1: 1, 2: 2, 3: 3}
	//only del key once, and not del the current iteration key
	var o sync.Once
	for i := range m {
		o.Do(func() {
			for _, key := range []int{1, 2, 3} {
				if key != i {
					fmt.Printf("when iteration key %d, del key %d\n", i, key)
					delete(m, key)
					break
				}
			}
		})
		fmt.Printf("%d%d ", i, m[i])
	}
}
